### The Graphics Pipeline
Most of the work of OpenGL is to convert 3D coordinates into colored 2D pixels into the screen.
This is done through what is called the "Graphics Pipeline". The pipeline can be divided into
two main parts. The first transforms the 3D coordinates into 2D coordinates and the second transforms the
2D coordinates into colored pixels.

The pipeline is composed of multiple steps that perform a specialized function (transformation)
over the data, and during the process small programs called "shaders" are sent to the GPU to be
processed. These shaders are written in [OpenGL Shading Language (GLSL)](https://www.khronos.org/opengl/wiki/Core_Language_(GLSL)).

Here is a high level graph of the pipeline:

<img src="https://fly.storage.tigris.dev/cmc-bucket/images/opengl_graphics_pipeline.jpg" alt="OpenGL Graphics Pipeline" width="600" />

The input to the pipeline is the `vertex data` which is a list of vertices. These vertices contain information in a 3D coordinate like the 
position x, y, and z. However, it can also store any other type of data.

**Vertex Shader:** This is the first part of the pipeline. The vertex shader takes as an input a single vertex and converts 3D coordinates into some other type
of 3D coordinates. During this process we can do some processing of the vertex attributes.

**Geometry Shader:** The output from the vertex shader is then passed to the geometry shader which takes a collection of vertices that represent a primitive.
The geometry shader converts the this primitive into a different primitive if needed by creating a different set of vertices.

**Primitive (Shape) Assembly:** The next step is the primitive or shape assembly, where the output from the vertex shader or the geometry shader is processed to assemble all the 
points from the primitive shape given.

**Rasterization Stage:** This process maps the primitives to the corresponding pixels in the final screen. The output of this process are fragments (a fragment in OpenGL is all the data required for OpenGL to render a single pixel) which are later used by the fragment shader. Before the fragments are passed to the fragment shader, a clipping process is performed to discard all the fragments that are not in view. This is done to improve the performance.

**Fragment Shader:** The fragment shader is what calculates the final color of a pixel in the screen. In this process all the information from the scene is processed
to calculate the final value of the color of each pixel. Here is where all the effects are applied, like the lights, the shadows and the colors.

### Vertex Input
All coordinates of a vertex need to be in 3D coordinates with x, y, and z. Additionally, these coordinates need to be normalized in a range
between `1.0` and `-1.0` for OpenGL to be able to process them. This means that a vertex is composed of x, y, and z components that are floats
between `1.0` and `-1.0`, where `x: 0.0, y: 0.0, z: 0.0` is the origin. Anything that extends the ranges will be discarted by OpenGL.

For example, if we want to create a triangle, we need 3 vertices:

```c
float vertices[] = {
    -0.5f, -0.5f, 0.0f,
    0.5f, -0.5f, 0.0f,
    0.0f, 0.5f, 0.0f
};
```

We want to create a 2D triangle and for that reason the z component of each vertex needs to be `0.0f`.

Now we can start the pipeline process for our triangle by sending our newly created *Vertex Input* to the 
first step of the process, the *Vertex Shader*. For this, we need to create some memory in the GPU where we
want to store the vertices. Then we need to configure how should OpenGL interpret the memory and how should the 
data be sent to the GPU.

The way to manage the memory is via the **vertex buffer object (VBO)** that can store a big amount of vertices
in the GPU to be processed. Setting as much data in the VBO as possible will make the processing very fast instead of
sending one vertex at a time.

Let us see how to generate a VBO:

```c
unsigned int VBO;
glGenBuffers(1, &VBO);  
```

OpenGL has multiple buffer types and we can use them at the same time as long as they have a different type.
For our triangle we will be using a `GL_ARRAY_BUFFER`. We can now bind the newly created VBO id to the 
`GL_ARRAY_BUFFER` with the `glBindBuffer` function:

```c
glBindBuffer(GL_ARRAY_BUFFER, VBO);
```

Now finally, we can copy the vextex data into the buffer memory:

```c
glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);
```

Here we need to specify how should the GPU manage the data by configuring it with either:

* `GL_STREAM_DRAW`: the data is set only once and used by the GPU at most a few times.
* `GL_STATIC_DRAW`: the data is set only once and used many times.
* `GL_DYNAMIC_DRAW`: the data is changed a lot and used many times.

For our triangle, the data does not change, and the data is used many times. Each render we will use
the same vertex data.

The next step is to create vertex and fragment shaders that will process this data.

### Vertex Shader
Vertex shaders are programmable and modern OpenGL requires that we provide at least one vertex shader
and one fragment shader. The shaders are written on a shader language called GLSL (OpenGL Shading Language),
which is very similar to C. Let us see an example:

```c
#version 330 core
layout (location = 0) in vec3 aPos;

void main()
{
    gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}
```

We need to specify the version of OpenGL that we are using and the type of profile that we are using, in 
this case `core` profile. In the previous example, we create a `vec3` named aPos and assign the input data
to it, in this case we only care about the position attribute.

Finally, the way to set the output of the shader is by assigning to a predefined `vec4` called `gl_Position`.
At the end of the main function, whatever is set on the `gl_Position` vector is going to be set as the output
of the shader.

Before we are able to use the shader we need to compile it.

```c
const char *vertexShaderSource = "#version 330 core\n"
"layout (location = 0) in vec3 aPos;\n"
"void main()\n"
"{\n"
"   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);\n"
"}\0";

unsigned int vertexShader;
vertexShader = glCreateShader(GL_VERTEX_SHADER);

glShaderSource(vertexShader, 1, &vertexShaderSource, NULL);
glCompileShader(vertexShader);

// Checking for compile time errors
int  success;
char infoLog[512];
glGetShaderiv(vertexShader, GL_COMPILE_STATUS, &success);
if(!success)
{
    glGetShaderInfoLog(vertexShader, 512, NULL, infoLog);
    std::cout << "ERROR::SHADER::VERTEX::COMPILATION_FAILED\n" << infoLog << std::endl;
}
```

If no errors are displayed after checking for compile time errors, means that our 
shader is compiled and ready to be used.

The next step is to create and compile the fragment shader which will give color to the pixels of
our triangle.

### Fragment Shader
The fragment shader will calculate the color output for the pixels of the triangle.

```c
#version 330 core
out vec4 FragColor;

void main()
{
    FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
```

The colors are in the form of RGBA (Red, Green, Blue, Alpha) and the values go from `0.0f` to `1.0f`
depending on the strength of each component. In this case, the color of the triangle with this fragment
will be in a orange tone. We also set an output variable called `FragColor` that is assigned and used as 
the output of the fragment.

The process for compiling the fragment is the following:

```c
const char *fragmentShaderSource = "#version 330 core\n"
"out vec4 FragColor;\n"
"void main()\n"
"{\n"
"   FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);\n"
"}\0";

unsigned int fragmentShader;
// Specify the type of shader to be GL_FRAGMENT_SHADER
fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);
glShaderSource(fragmentShader, 1, &fragmentShaderSource, NULL);
glCompileShader(fragmentShader);

// Do not forget to check for compile time errors
int  success;
char infolog[512];
glGetShaderiv(fragmentShader, GL_COMPILE_STATUS, &success);
if(!success)
{
    glGetShaderInfoLog(fragmentShader, 512, NULL, infoLog);
    std::cout << "ERROR::SHADER::FRAGMENT::COMPILATION_FAILED\n" << infoLog << std::endl;
}
```

Now that we have both the vertex and fragment shader compiled, the only thing that is there to do
is to link both with a *Shader Program*

### Shader Program
The shader program object combines multiple shaders by linking the output of one shader to the input of 
the next shader. For rendering objects, the shader program object must be activated.

Let us create our first program:

```c
unsigned int shaderProgram;
shaderProgram = glCreateProgram();

// The order of attaching matters
glAttachShader(shaderProgram, vertexShader);
glAttachShader(shaderProgram, fragmentShader);
// Link the entire program
glLinkProgram(shaderProgram);

// Check for linking errors
int  success;
char infolog[512];
glGetProgramiv(shaderProgram, GL_LINK_STATUS, &success);
if(!success) {
    glGetProgramInfoLog(shaderProgram, 512, NULL, infoLog);
    std::cout << "ERROR::SHADER::PROGRAM::LINKING_FAILED\n" << infoLog << std::endl;
}

// Use the shader program
glUseProgram(shaderProgram);

// Finally, delete the vertex and fragment shader objects. We don't need them anymore
glDeleteShader(vertexShader);
glDeleteShader(fragmentShader);
```

### Linking Vertex Attributes
So far, we can specify any type of vertex attributes for our vertex shader, but OpenGL does 
not know how to interpret those attributes. We can tell OpenGL how to interpret the vertex attributes
by specifing how the memory is layed out for each of the attributes.

We can do this by using the `glVertexAttribPointer` function.

```c
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
glEnableVertexAttribArray(0); 
```

Here we tell OpenGL that:
* The first argument is which vertex attribute we are configuring. In this case is `0` because we specified in
the shader that the location of the position attribute was `0` in the line `layout (location = 0)`.
* The second argument is the size of the vertex attribute. The position attribute is a `vec3` so it has 3 components.
* The third argument is the type of the attribute components. In OpenGL `vec*` are floats.
* The fourth argument tells if the data should be normalized.
* The fifth argument is the stride or how far is the next set of vertex attribute. This means that for the first vertex,
the position starts at the 0 index and the position for the next vertex is 3 floats away.
* The last argument is the offset of where the position attribute starts. It is a void pointer so we need to cast it.

By default, vertex attributes are disabled this is why we need to call `glEnableVertexAttribArray(0)` specifing the 
vertex attribute location as an argument.

Every time we want to draw an object we need to make the same calls:

```c
// 0. copy our vertices array in a buffer for OpenGL to use
glBindBuffer(GL_ARRAY_BUFFER, VBO);
glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);
// 1. then set the vertex attributes pointers
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
glEnableVertexAttribArray(0);
// 2. use our shader program when we want to render an object
glUseProgram(shaderProgram);
// 3. now draw the object 
someOpenGLFunctionThatDrawsOurTriangle();
```

This could become very repetitive when we are trying to render lots of objects with many vertex attributes.
For this reason there is something called the **Vertex Array Object (VAO)** which stores all the state configurations
of objects to be used later.

### Vertex Array Object

A vertex array object (VAO) is bound like a vertex buffer object and all subsequent calls to configure the attribute pointer
will be stored in the array. Later, whenever we want to draw some vertices, we just need to specify the VAO. If we want to
draw a different type of object, we just bind the VAO where that object was configured and make the draw calls.

The process to create a VAO is the following:

```c
GLuint VAO;
glGenVertexArrays(1, &VAO);

// We run this initialization code once unless the vertices change often
// 1. bind Vertex Array Object
glBindVertexArray(VAO);
// 2. copy our vertices array in a buffer for OpenGL to use
glBindBuffer(GL_ARRAY_BUFFER, VBO);
glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);
// 3. then set our vertex attributes pointers
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
glEnableVertexAttribArray(0);

// 4. draw the object
glUseProgram(shaderProgram);
glBindVertexArray(VAO);
someOpenGLFunctionThatDrawsOurTriangle();
```

By binding the new VAO and then binding the VBO, configuring the vertex attributes pointers
and enabling the attributes will be stored in the currently bound VAO.

Then, to draw a specific VBO, we just need to bind the VAO before calling the draw functions.

### The Triangle

Finally, we have everything we need to paint the triangle. For this we can use the `glDrawArrays`
function passing the `GL_TRIANGLES` option

```c
// ...
glUseProgram(shaderProgram);
glBindVertexArray(VAO);
glDrawArrays(GL_TRIANGLES, 0, 3);
```

We need to specify the type of the primitive that we want to draw, the starting index of the vertex
array that we would like to draw (`0`), and the number of vertices that we want to draw (`3`).

<img src="https://fly.storage.tigris.dev/cmc-bucket/images/opengl_triangle.png" alt="OpenGL Orange Triangle" width="600" />

Finally, we got the triangle in the screen, Yay!.

### Element Buffer Object 

There is one more topic to cover. Imagine that now we want to paint a rectangle. In OpenGL we need 
to use the primitives to compose other shapes, so in this case we need to paint two triangles next 
to each other. For this we will need to instruct OpenGL how to connect each vertex with something called
**Element Buffer Object (EBO)**.

Let us declare our new vertices:

```c
float vertices[] = {
    // first triangle
     0.5f,  0.5f, 0.0f,  // top right
     0.5f, -0.5f, 0.0f,  // bottom right
    -0.5f,  0.5f, 0.0f,  // top left 
    // second triangle
     0.5f, -0.5f, 0.0f,  // bottom right
    -0.5f, -0.5f, 0.0f,  // bottom left
    -0.5f,  0.5f, 0.0f   // top left
};
```
As you can see, there are two vertices repeated the `bottom right` and the `top left`. This is very inefficient.
We can do something smarter. We can just specify 4 unique vertices and instruct OpenGL the order in which we want
to paint those vertices to build the shape that we want.

Element Buffer Objects can specify a list of vertex indices to draw in order.

Let us see how this works:

```c
// First specify the unique vertices that we need for our rectangle
float vertices[] = {
     0.5f,  0.5f, 0.0f,  // top right
     0.5f, -0.5f, 0.0f,  // bottom right
    -0.5f, -0.5f, 0.0f,  // bottom left
    -0.5f,  0.5f, 0.0f   // top left 
};

// The indices construct the triangles reusing the vertex that we have defined
GLuint indices[] = {
    0, 1, 3,   // first triangle
    1, 2, 3    // second triangle
};


// Similar to the VBO, we create a new EBO, bind it and copy the indices data to it.
unsigned int EBO;
glGenBuffers(1, &EBO);

glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, EBO);
glBufferData(GL_ELEMENT_ARRAY_BUFFER, sizeof(indices), indices, GL_STATIC_DRAW);

// Finally, instead of drawing arrays, we instruct OpenGL to draw elements
glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, EBO);
glDrawElements(GL_TRIANGLES, 6, GL_UNSIGNED_INT, 0);
```

The VAO can also store the last bound EBO. This way we can remember the EBO and 
draw it by binding a VAO.

```c
// Bind Vertex Array Object
glBindVertexArray(VAO);

// Copy our vertices array in a vertex buffer for OpenGL to use
glBindBuffer(GL_ARRAY_BUFFER, VBO);
glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);

// Copy our index array in a element buffer for OpenGL to use
glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, EBO);
glBufferData(GL_ELEMENT_ARRAY_BUFFER, sizeof(indices), indices, GL_STATIC_DRAW);

// Then set the vertex attributes pointers
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
glEnableVertexAttribArray(0);

// Drawing code
glUseProgram(shaderProgram);
glBindVertexArray(VAO);
glDrawElements(GL_TRIANGLES, 6, GL_UNSIGNED_INT, 0);
glBindVertexArray(0);
```

After compiling the program and running it you should see the rectangle!

<img src="https://fly.storage.tigris.dev/cmc-bucket/images/opengl_rectangle.png" alt="OpenGL Orange Rectangle" width="600" />

In the next guides we will dive deep and explore some of these concepts.
