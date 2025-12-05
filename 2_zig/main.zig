const std = @import("std");

const input = @embedFile("input.txt");

pub fn main() !void {
    var timer = try std.time.Timer.start();
    try part1();
    var elapsed = timer.read();
    std.debug.print("elapsed: {}ms\n", .{elapsed / 1_000_000});

    try part2();
    elapsed = timer.read();
    std.debug.print("elapsed: {}ms\n", .{elapsed / 1_000_000});
}

// just iterates through the range, and then compares the first and second half of even digit numbers
fn part1() !void {
    var sum: usize = 0;
    var ranges_iter = get_ranges();
    while (ranges_iter.next()) |range| {
        var range_iter = std.mem.tokenizeAny(u8, range, "-");
        var from = try std.fmt.parseInt(usize, range_iter.next().?, 10);
        const to = try std.fmt.parseInt(usize, range_iter.next().?, 10);
        while (from < to) {
            from += 1;
            const num_of_digits = get_num_of_digits(from);

            if (num_of_digits % 2 == 0) {
                const fh = get_first_half(from, num_of_digits);
                const sh = get_second_half(from, num_of_digits);
                if (fh == sh) {
                    sum += from;
                }
            }
        }
    }

    std.debug.print("part 1: {d}\n", .{sum});
}

// my part 2 converts each num to string to make comparisons / checks for repetition easier
fn part2() !void {
    var sum: usize = 0;

    // buffer for numbers that we stringify in the loop
    var buf: [32]u8 = undefined;

    var ranges_iter = get_ranges();
    while (ranges_iter.next()) |range| {
        var range_iter = std.mem.tokenizeAny(u8, range, "-");
        var from = try std.fmt.parseInt(usize, range_iter.next().?, 10);
        const to = try std.fmt.parseInt(usize, range_iter.next().?, 10);
        while (from <= to) {
            var is_repeating = false;

            // converting num to string
            const slice = try std.fmt.bufPrint(&buf, "{d}", .{from});
            const num_of_digits = slice.len;
            var chunk_size = num_of_digits / 2;

            while (chunk_size > 0) {
                if (num_of_digits % chunk_size > 0) {
                    chunk_size -= 1;
                    continue;
                }

                if (has_identical_chunks(slice, chunk_size)) {
                    is_repeating = true;
                    break;
                }

                chunk_size -= 1;
            }

            if (is_repeating) {
                sum += from;
            }

            from += 1;
        }
    }

    std.debug.print("part 2: {d}\n", .{sum});
}

fn get_ranges() std.mem.TokenIterator(u8, .any) {
    var lines = std.mem.tokenizeAny(u8, input, "\r\n");
    const first_line = lines.next().?;
    const ranges_iter = std.mem.tokenizeAny(u8, first_line, ",");
    return ranges_iter;
}

fn get_num_of_digits(n: usize) usize {
    var count: usize = 0;
    var d = n;
    while (d > 0) {
        d = @divTrunc(d, 10);
        count += 1;
    }
    return count;
}

fn get_first_half(n: usize, len: usize) usize {
    const result = @divTrunc(n, pow(10, len / 2));
    return result;
}

fn get_second_half(n: usize, len: usize) usize {
    const result = @mod(n, pow(10, len / 2));
    return result;
}

fn has_identical_chunks(str: []const u8, chunk_size: usize) bool {
    const initial_chunk = str[0..chunk_size];
    const iters = str.len / chunk_size;
    for (1..iters) |i| {
        const next_chunk_idx = i * chunk_size;
        const next_chunk = str[next_chunk_idx .. next_chunk_idx + chunk_size];
        if (!std.mem.eql(u8, initial_chunk, next_chunk)) {
            return false;
        }
    }

    return true;
}

fn pow(n: usize, x: usize) usize {
    if (x == 1) {
        return n;
    }

    return n * pow(n, x - 1);
}

test "get_num_of_digits" {
    try std.testing.expectEqual(3, get_num_of_digits(123));
}

test "get_first_half" {
    try std.testing.expectEqual(1, get_first_half(12, 2));
    try std.testing.expectEqual(10, get_first_half(1010, 4));
    try std.testing.expectEqual(123, get_first_half(123456, 6));
}

test "get_second_half" {
    try std.testing.expectEqual(2, get_second_half(12, 2));
    try std.testing.expectEqual(10, get_second_half(1010, 4));
    try std.testing.expectEqual(456, get_second_half(123456, 6));
}

test "has_identical_chunks" {
    try std.testing.expect(has_identical_chunks("121212", 2));
    try std.testing.expect(!has_identical_chunks("1213", 2));
    try std.testing.expect(has_identical_chunks("111", 1));
}
