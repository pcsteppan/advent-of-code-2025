const std = @import("std");
const input = @embedFile("input.txt");
const line_length = 100;

pub fn main() !void {
    var sum1: usize = 0;
    var sum2: usize = 0;
    var lines = std.mem.tokenizeAny(u8, input, "\r\n");
    while (lines.next()) |line| {
        sum1 += try get_max_n_jolts(line, 2);
        sum2 += try get_max_n_jolts(line, 12);
    }

    std.debug.print("part 1: {d}\n", .{sum1});
    std.debug.print("part 2: {d}\n", .{sum2});
}

// function previously used for part 1
// made more generic in part 2 with the get_max_n_jolts fn
fn get_max_jolts(line: []const u8) usize {
    var n1: usize = 0;
    var n2: usize = 0;

    for (0..line_length) |i| {
        const n = line[i] - '0';

        if (n > n1) {
            if (i < line_length - 1) {
                n1 = n;
                n2 = 0;
            } else {
                n2 = n;
            }
        } else if (n > n2) {
            n2 = n;
        }
    }

    return n1 * 10 + n2;
}

const ArgumentError = error{ExceedsMaxLength};

fn get_max_n_jolts(line: []const u8, n: usize) !usize {
    if (n > line_length) {
        return ArgumentError.ExceedsMaxLength;
    }

    var buf: [line_length]u8 = .{0} ** line_length;
    const maxes = buf[0..n];
    const start_index_offset = line_length - n;

    for (0..line_length) |i| {
        const jolt = line[i] - '0';
        // -| is a 'saturating subtraction' operator that clamps the lower bound at 0
        const start = i -| start_index_offset;
        var clear_remaining = false;
        for (start..n) |j| {
            if (!clear_remaining) {
                const max = &maxes[j];
                if (jolt > max.*) {
                    max.* = jolt;
                    clear_remaining = true;
                }
            } else {
                // clear rest of buffer if we found a new max earlier in the bank
                maxes[j] = 0;
            }
        }
    }

    return sum_digit_array(maxes);
}

fn sum_digit_array(digits: []const u8) usize {
    var sum: usize = 0;
    for (digits) |digit| {
        sum *= 10;
        sum += digit;
    }
    return sum;
}

test "get_max_jolts" {
    try std.testing.expectEqual(99, get_max_jolts("8979" ** 25)); // repeat 25 times to get to 100 line length
}

test "get_max_n_jolts" {
    try std.testing.expectEqual(999, get_max_n_jolts("8979" ** 25, 3));
}
