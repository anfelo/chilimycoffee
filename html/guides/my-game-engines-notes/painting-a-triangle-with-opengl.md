### The Graphics Pipeline
Most of the work of OpenGL is to convert 3D coordinates into colored 2D pixels in the screen.
This is done through what is called the "Graphics Pipeline". The pipeline can be divided into
two main parts. The first transforms the 3D coordinates into 2D coordinates and the second transforms the
2D coordinates into colored pixels.

The pipeline is composed of multiple steps that perform a specialized function (transformation)
over the data, and during the process small programs called "shaders" are sent to the GPU to be
processed. These shaders are written in [OpenGL Shading Language (GLSL)](https://www.khronos.org/opengl/wiki/Core_Language_(GLSL)).

Here is a high level graph of the pipeline:

<img src="https://fly.storage.tigris.dev/cmc-bucket/images/opengl_graphics_pipeline.jpg" alt="OpenGL Graphics Pipeline" width="600" />

The input to the pipeline is the `vertex data` which is a list of vertices. These vertices contain information in a 3D coordinate like the 
position x, y, and z. However, it can also store any type of data.

### 
