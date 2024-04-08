

shapes_rectangle_scaling: ./src/examples/shapes/shapes_rectangle_scaling/shapes_rectangle_scaling.go c_raylib
	go build -o ./bin/examples/shapes/shapes_rectangle_scaling ./src/examples/shapes/shapes_rectangle_scaling/shapes_rectangle_scaling.go

shapes_bouncing_ball: ./src/examples/shapes/shapes_bouncing_ball/shapes_bouncing_ball.go c_raylib
	go build -o ./bin/examples/shapes/shapes_bouncing_ball ./src/examples/shapes/shapes_bouncing_ball/shapes_bouncing_ball.go

shapes_basic_shapes: ./src/examples/shapes/shapes_basic_shapes/shapes_basic_shapes.go c_raylib
	go build -o ./bin/examples/shapes/shapes_basic_shapes ./src/examples/shapes/shapes_basic_shapes/shapes_basic_shapes.go

shapes_collision_area: ./src/examples/shapes/shapes_collision_area/shapes_collision_area.go c_raylib
	go build -o ./bin/examples/shapes/shapes_collision_area ./src/examples/shapes/shapes_collision_area/shapes_collision_area.go

shapes_colors_palette: ./src/examples/shapes/shapes_colors_palette/shapes_colors_palette.go c_raylib
	go build -o ./bin/examples/shapes/shapes_colors_palette ./src/examples/shapes/shapes_colors_palette/shapes_colors_palette.go

shapes_draw_circle_sector: ./src/examples/shapes/shapes_draw_circle_sector/shapes_draw_circle_sector.go c_raylib
	go build -o ./bin/examples/shapes/shapes_draw_circle_sector ./src/examples/shapes/shapes_draw_circle_sector/shapes_draw_circle_sector.go

shapes_draw_rectangle_rounded: ./src/examples/shapes/shapes_draw_rectangle_rounded/shapes_draw_rectangle_rounded.go c_raylib
	go build -o ./bin/examples/shapes/shapes_draw_rectangle_rounded ./src/examples/shapes/shapes_draw_rectangle_rounded/shapes_draw_rectangle_rounded.go


shapes: shapes_rectangle_scaling shapes_basic_shapes shapes_bouncing_ball shapes_collision_area shapes_colors_palette shapes_draw_circle_sector shapes_draw_rectangle_rounded

c_raylib: $(shell find raylib/src -type f \( -iname \*.c -o -iname \*.h \))
	cd ./raylib/src ;\
	make PLATFORM=PLATFORM_DESKTOP ;\
	cd ../..

examples: shapes

all: examples