const std = @import("std");
const input = @embedFile("input.txt");
const max_width = 136;
const max_height = 136;

pub fn main() !void {
    // using split rather than tokenize to keep the empty line that divides the input sections
    var lines = std.mem.splitScalar(u8, input, '\n');
    const data = try Data.init(&lines);

    var sum: usize = 0;
    for (data.samples) |sample| {
        for (data.ranges) |range| {
            if (range[0] < sample and range[1] > sample) {
                sum += 1;
                break;
            }
        }
    }

    std.debug.print("part 1: {d}\n", .{sum});
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
            // split
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
