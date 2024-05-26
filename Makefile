

shapes_draw_ring: ./src/examples/shapes/shapes_draw_ring/shapes_draw_ring.go c_raylib
	go build -o ./bin/examples/shapes/shapes_draw_ring ./src/examples/shapes/shapes_draw_ring/shapes_draw_ring.go

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

shapes_easings_ball_anim: ./src/examples/shapes/shapes_easings_ball_anim/shapes_easings_ball_anim.go c_raylib
	go build -o ./bin/examples/shapes/shapes_easings_ball_anim ./src/examples/shapes/shapes_easings_ball_anim/shapes_easings_ball_anim.go

shapes_easings_box_anim: ./src/examples/shapes/shapes_easings_box_anim/shapes_easings_box_anim.go c_raylib
	go build -o ./bin/examples/shapes/shapes_easings_box_anim ./src/examples/shapes/shapes_easings_box_anim/shapes_easings_box_anim.go

shapes_easings_rectangle_array: ./src/examples/shapes/shapes_easings_rectangle_array/shapes_easings_rectangle_array.go c_raylib
	go build -o ./bin/examples/shapes/shapes_easings_rectangle_array ./src/examples/shapes/shapes_easings_rectangle_array/shapes_easings_rectangle_array.go

shapes_following_eyes: ./src/examples/shapes/shapes_following_eyes/shapes_following_eyes.go c_raylib
	go build -o ./bin/examples/shapes/shapes_following_eyes ./src/examples/shapes/shapes_following_eyes/shapes_following_eyes.go

shapes_logo_raylib: ./src/examples/shapes/shapes_logo_raylib/shapes_logo_raylib.go
	go build -o ./bin/examples/shapes/shapes_logo_raylib ./src/examples/shapes/shapes_logo_raylib/shapes_logo_raylib.go

shapes_lines_bezier: ./src/examples/shapes/shapes_lines_bezier/shapes_lines_bezier.go
	go build -o ./bin/examples/shapes/shapes_lines_bezier ./src/examples/shapes/shapes_lines_bezier/shapes_lines_bezier.go

shapes_logo_raylib_anim: ./src/examples/shapes/shapes_logo_raylib_anim/shapes_logo_raylib_anim.go
	go build -o ./bin/examples/shapes/shapes_logo_raylib_anim ./src/examples/shapes/shapes_logo_raylib_anim/shapes_logo_raylib_anim.go

shapes: shapes_draw_ring shapes_basic_shapes shapes_bouncing_ball shapes_collision_area shapes_colors_palette shapes_draw_circle_sector shapes_draw_rectangle_rounded shapes_easings_ball_anim shapes_easings_box_anim shapes_easings_rectangle_array shapes_following_eyes shapes_rectangle_scaling shapes_logo_raylib shapes_lines_bezier shapes_logo_raylib_anim

c_raylib: $(shell find raylib/src -type f \( -iname \*.c -o -iname \*.h \))
	cd ./raylib/src ;\
	make PLATFORM=PLATFORM_DESKTOP ;\
	cd ../..

examples: shapes

all: examples

rm_go_ralib_bin: 
	if [ -d "bin" ]; then rm -r ./bin; fi

clean_raylib:
	cd ./raylib/src ;\
	make clean ;\
	cd ../..


clean: rm_go_ralib_bin clean_raylib