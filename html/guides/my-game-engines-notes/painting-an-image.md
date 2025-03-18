The main idea of this article is to build on top of the previous articles and paint
an image, also to be able to move it around, scale it, and rotate it. But first we need
to explain some basic topics that we will be using. Let us start with the shaders:

### What are shaders

Shaders are little programs running in the GPU. They process inputs and generate
outputs that can be used by other shader programs.

Shaders are written on a graphics programing language called GLSL (OpenGL Shading Language)
that is similar to the C language. Each shader has a main function (the entry point) where 
the inputs are processed and any outputs are calculated.

In the previous article we saw an example of a vertex shader and a fragment shader
they all look similar but they run in different stages of the graphics pipeline.

The next one is a simple example of a vertex shader:

```c
#version 330 core
layout (location = 0) in vec3 a_pos; // the position variable has attribute position 0

out vec4 vertex_color; // specify a color output to the fragment shader

void main()
{
    gl_Position = vec4(a_pos, 1.0); // see how we directly give a vec3 to vec4's constructor
    vertex_color = vec4(0.5, 0.0, 0.0, 1.0); // set the output variable to a dark-red color
}
```

First we specify the OpenGL version (v3.3 core) that it supports. Then we declare an
input for the position `a_pos` of type `vec3` and we specify the attribute location (this
will be used in the program to locate the attribute). We also declare a vertex color output
`vertex_color` that we can set. By setting this output, it will be available to the next
shaders in the pipeline to do any computation. The last thing that we see here is 
that we set the `gl_Position` variable. This is a special variable that OpenGL
uses to set the position of a vertex.

The following is an example of the fragment shader:

```c
#version 330 core
out vec4 FragColor;

in vec4 vertex_color; // the input variable from the vertex shader (same name and same type)

void main()
{
    FragColor = vertex_color;
}
```

The fragment shader's main job is to set the pixel color that will be displayed.
We only need to set `FragColor` which is a special output that OpenGL will use to
know what is the color of the following output.
In theory it is a bit more complex than that, but for now this should be fine to 
continue.

### Shaders Inputs

There are 3 main types of inputs a shader can accept, and they are used for different
purposes. But in general, inputs are used to pass information to a shader where they are
used to calculate either vertex positions (vertex shader) or pixel color of a fragment 
(fragment shader).

Let us start with Attributes

#### Attributes

Attributes are set in our program code and are passed to the vertex shader. These 
are usually the values of the vertices positions, normals and texture coordinates.

We saw this one in the vertex shader above where we specify an attribute called 
`a_pos` which contains the information of the position of the vertices.

```glsl
...
layout (location = 0) in vec3 a_pos;
...
```
Here we use the keyword `layout` plus the location of the attribute to tell OpenGL
which slot this attribute is in. Here we also use the `in` keyword to specify that
this variable is an input and finally we specify the type of the attribute. In this
case we have a vector with 3 components for the vertex position in x, y and z.

Attibutes are set in code in this way:

```cpp
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
glEnableVertexAttribArray(0);
```

We use the `glVertexAttribPointer` function to specify the layout of the vertex data.
The first argument is the location of the attribute (same location as defined in the
shader). Then we define the number of components of this attribute (`a_pos` is a vec3 so 
we need 3). Then comes the type of the data (floats) and whether the data should be 
normalized (`GL_FALSE`). Then we set the offset of each vertex data, in this case we only
have positions in the array of data so the next vertex position would be 3 floats away,
is distance (this is also called the stride). Finally we set the offset pointer to
this attribute. Here we only have one attribute (`vec3 a_pos`) but this when we have 
more attributes packed inside a vertex array we need to make some calculation of 
where they are.

Finally, we enable the vertex attibute using the `glEnableVertexAttribArray` and passing
the index or location specified in the previous step.
