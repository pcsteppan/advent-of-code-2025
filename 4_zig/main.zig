const std = @import("std");
const input = @embedFile("input.txt");
const max_width = 136;
const max_height = 136;

pub fn main() !void {
    var lines = std.mem.tokenizeAny(u8, input, "\r\n");
    const grid = Grid.init(&lines);

    var result = grid.remove_and_count();
    var total_removed: usize = result.count;
    std.debug.print("part 1: {d}\n", .{total_removed});
    while (result.count > 0) {
        result = result.grid.remove_and_count();
        total_removed += result.count;
    }
    std.debug.print("part 1: {d}\n", .{total_removed});
}

const Grid = struct {
    width: usize,
    height: usize,
    buffer: [max_height][max_width]u8,

    pub fn init(lines: *std.mem.TokenIterator(u8, .any)) Grid {
        var height: usize = 0;
        var width: usize = 0;
        var buffer: [max_height][max_width]u8 = undefined;
        while (lines.next()) |line| {
            @memcpy(buffer[height][0..line.len], line);
            height += 1;
            width = line.len;
        }

        return .{
            .width = width,
            .height = height,
            .buffer = buffer,
        };
    }

    pub fn get(self: *const Grid, x: isize, y: isize) ?u8 {
        if (x < 0 or y < 0 or x >= self.width or y >= self.height) {
            return null;
        }

        return self.buffer[@intCast(y)][@intCast(x)];
    }

    pub fn remove_and_count(self: *const Grid) struct { grid: Grid, count: usize } {
        // assignments of structs are by value, so this copies the grid
        var new_grid = self.*;

        // can't use negative numbers in ranges, e.g, -1..2,
        // so using this method instead
        const kernel = [3]isize{ -1, 0, 1 };
        var sum: usize = 0;

        for (0..self.height) |y| {
            for (0..self.width) |x| {
                const is_at = self.get(@intCast(x), @intCast(y)).? == '@';
                if (!is_at) {
                    continue;
                }

                var neighbor_count: usize = 0;
                for (kernel) |dx| {
                    for (kernel) |dy| {
                        const on_self = dx == 0 and dy == 0;
                        if (on_self) {
                            continue;
                        }
                        if (self.get(@as(isize, @intCast(x)) + dx, @as(isize, @intCast(y)) + dy)) |c| {
                            if (c == '@') {
                                neighbor_count += 1;
                            }
                        }
                    }
                }

                if (neighbor_count < 4) {
                    new_grid.buffer[y][x] = '.';
                    sum += 1;
                }
            }
        }

        return .{ .grid = new_grid, .count = sum };
    }
};
