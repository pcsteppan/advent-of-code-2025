const std = @import("std");
const input = @embedFile("input.txt");

pub fn main() !void {
    // using split rather than tokenize to keep the empty line that divides the input sections
    var lines = std.mem.splitScalar(u8, input, '\n');
    const data = try Data.init(&lines);

    var sum: usize = 0;
    for (data.samples[0..data.sample_count]) |sample| {
        for (data.ranges[0..data.range_count]) |range| {
            if (range[0] < sample and range[1] > sample) {
                sum += 1;
                break;
            }
        }
    }

    std.debug.print("part 1: {d}\n", .{sum});

    const range_array: RangeArray = try RangeArray.init(&data.ranges, data.range_count);
    std.debug.print("part 2: {d}", .{range_array.find_real_numbers_in_ranges()});
}

const max_ranges = 200;
const max_samples = 1001;
const Data = struct {
    range_count: usize,
    sample_count: usize,
    ranges: [max_ranges][2]usize,
    samples: [max_samples]usize,

    pub fn init(lines: *std.mem.SplitIterator(u8, .scalar)) !Data {
        var range_count: usize = 0;
        var sample_count: usize = 0;
        var ranges_buffer: [max_ranges][2]usize = undefined;
        var samples: [max_samples]usize = undefined;

        while (lines.next()) |line| {
            const trimmed = std.mem.trimRight(u8, line, "\r\n");
            if (trimmed.len == 0) {
                break;
            }

            var ab = std.mem.tokenizeScalar(u8, trimmed, '-');
            const a = try std.fmt.parseInt(usize, ab.next().?, 10);
            const b = try std.fmt.parseInt(usize, ab.next().?, 10);
            ranges_buffer[range_count][0] = a;
            ranges_buffer[range_count][1] = b;

            range_count += 1;
        }

        while (lines.next()) |line| {
            const trimmed = std.mem.trimRight(u8, line, "\r\n");
            if (trimmed.len == 0) {
                // final line
                break;
            }
            samples[sample_count] = try std.fmt.parseInt(usize, trimmed, 10);

            sample_count += 1;
        }

        return .{ .range_count = range_count, .sample_count = sample_count, .samples = samples, .ranges = ranges_buffer };
    }
};

// Data structure which maintains two important invariants
// that ranges are exclusive / do not overlap with each other
// and they're kept in ascending order
const RangeArray = struct {
    range_count: usize,
    ranges: [2][max_ranges]usize,

    pub fn init(ranges: *const [max_ranges][2]usize, count: usize) !RangeArray {
        var self: RangeArray = .{
            .range_count = 0,
            .ranges = undefined,
        };

        for (0..count) |i| {
            try self.add(ranges[i][0], ranges[i][1]);
        }

        return self;
    }

    // mutates data buffers directly
    // finds where to add the new range, updates, and shifts buffers accordingly
    fn add(self: *RangeArray, a: usize, b: usize) !void {
        // find first range which contains a
        const start_idx, const start_intersects = self.find_index_to_insert_at(a);

        // add to end
        if (start_idx == self.range_count) {
            try self.insert(self.range_count, a, b);
            return;
        }

        const end_idx, const end_intersects = self.find_index_to_insert_at(b);

        var ranges_intersected: usize = end_idx - start_idx;
        if (end_intersects) ranges_intersected += 1;

        const existing_start = if (start_intersects) self.ranges[0][start_idx] else std.math.maxInt(usize);
        const start = @min(existing_start, a);

        const existing_end = if (end_intersects) self.ranges[1][end_idx] else std.math.minInt(usize);
        const end = @max(existing_end, b);

        try self.remove(start_idx, ranges_intersected);
        try self.insert(start_idx, start, end);
    }

    // returns index and if it intersects the range there (otherwise it would mean it goes before that range)
    fn find_index_to_insert_at(self: *const RangeArray, n: usize) struct { usize, bool } {
        // no data yet, so we'd insert at the start
        if (self.range_count == 0) {
            return .{ 0, false };
        }

        var i: usize = 0;

        while (i < self.range_count) {
            const a = self.ranges[0][i];
            const b = self.ranges[1][i];
            if (n < a and n < b) {
                return .{ i, false };
            } else if (n >= a and n <= b) {
                return .{ i, true };
            }

            i += 1;
        }

        // the num is greater than all existing ranges
        return .{ self.range_count, false };
    }

    fn remove(self: *RangeArray, index: usize, count: usize) !void {
        if (count == 0) {
            return;
        }

        if (index > self.range_count or index + count > self.range_count) {
            return error.Error;
        }

        for (index..self.range_count - count) |i| {
            self.ranges[0][i] = self.ranges[0][i + count];
            self.ranges[1][i] = self.ranges[1][i + count];
        }

        self.range_count -= count;
    }

    // adds new range into buffers, and shifts remaining data accordingly
    fn insert(self: *RangeArray, index: usize, a: usize, b: usize) !void {
        if (index > self.range_count) {
            return error.Error;
        }

        var i = self.range_count;
        while (i > index) {
            self.ranges[0][i] = self.ranges[0][i - 1];
            self.ranges[1][i] = self.ranges[1][i - 1];
            i -= 1;
        }

        self.ranges[0][index] = a;
        self.ranges[1][index] = b;
        self.range_count += 1;
    }

    pub fn find_real_numbers_in_ranges(self: *const RangeArray) usize {
        var sum: usize = 0;
        for (0..self.range_count) |i| {
            sum += self.ranges[1][i] - self.ranges[0][i] + 1;
        }
        return sum;
    }

    pub fn format(
        self: *const RangeArray,
        writer: anytype,
    ) !void {
        try writer.print("RangeArray({} ranges): ", .{self.range_count});
        for (0..self.range_count) |i| {
            if (i > 0) try writer.writeAll(", ");
            try writer.print("{}-{}", .{ self.ranges[0][i], self.ranges[1][i] });
        }
    }
};
