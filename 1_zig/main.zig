const std = @import("std");
const input = @embedFile("input.txt");

pub fn main() !void {
    var position: isize = 50;
    var zero_count: usize = 0;
    var click_count: usize = 0;

    var lines_iter = std.mem.tokenizeAny(u8, input, "\r\n");
    while (lines_iter.next()) |line| {
        var delta: isize = try std.fmt.parseInt(isize, line[1..], 10);

        const isNegative = line[0] == 'L';
        if (isNegative) {
            delta *= -1;
        }

        const clicks, const pos = count_clicks(position, delta);
        // part 2
        click_count += clicks;
        position = pos;

        // part 1
        if (position == 0) {
            zero_count += 1;
        }
    }

    std.debug.print("part 1: zero count {d}\n", .{zero_count});
    std.debug.print("part 2: click count {d}\n", .{click_count});
}

fn count_clicks(from: isize, delta: isize) struct { usize, isize } {
    // delta of 0 is not allowed / expected in the input
    var new_pos = from;
    var hits: usize = 0;
    const step: isize = if (delta > 0) 1 else -1;
    var ticks_remaining: usize = @abs(delta);

    while (ticks_remaining > 0) {
        ticks_remaining -= 1;
        new_pos += step;
        new_pos = @mod(new_pos, 100);
        if (new_pos == 0) {
            hits += 1;
        }
    }

    return .{ hits, new_pos };
}

test "count_clicks" {
    // During the rotation
    try std.testing.expectEqual(.{ 1, 1 }, count_clicks(99, 2));
    try std.testing.expectEqual(.{ 1, 0 }, count_clicks(99, 1));
    try std.testing.expectEqual(.{ 0, 1 }, count_clicks(0, 1));
    try std.testing.expectEqual(.{ 1, 0 }, count_clicks(0, 100));
    try std.testing.expectEqual(.{ 1, 0 }, count_clicks(0, -100));
    try std.testing.expectEqual(.{ 2, 50 }, count_clicks(0, 250));
    try std.testing.expectEqual(.{ 3, 0 }, count_clicks(50, 250));
    // At the end of the rotation
    try std.testing.expectEqual(.{ 1, 0 }, count_clicks(1, -1));
    try std.testing.expectEqual(.{ 2, 0 }, count_clicks(1, -101));
    try std.testing.expectEqual(.{ 2, 0 }, count_clicks(99, 101));
}
