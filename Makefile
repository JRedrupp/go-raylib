

shapes_rectangle_scaling: ./src/examples/shapes/shapes_rectangle_scaling/shapes_rectangle_scaling.go c_raylib
	go build -o ./bin/examples/shapes/shapes_rectangle_scaling ./src/examples/shapes/shapes_rectangle_scaling/shapes_rectangle_scaling.go

shapes_bouncing_ball: ./src/examples/shapes/shapes_bouncing_ball/shapes_bouncing_ball.go c_raylib
	go build -o ./bin/examples/shapes/shapes_bouncing_ball ./src/examples/shapes/shapes_bouncing_ball/shapes_bouncing_ball.go

shapes_basic_shapes: ./src/examples/shapes/shapes_basic_shapes/shapes_basic_shapes.go c_raylib
	go build -o ./bin/examples/shapes/shapes_basic_shapes ./src/examples/shapes/shapes_basic_shapes/shapes_basic_shapes.go

shapes: shapes_rectangle_scaling shapes_basic_shapes shapes_bouncing_ball

c_raylib: $(shell find raylib/src -type f \( -iname \*.c -o -iname \*.h \))
	cd ./raylib/src ;\
	make PLATFORM=PLATFORM_DESKTOP ;\
	cd ../..

examples: shapes

all: examples