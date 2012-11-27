package glutil

import gl "code.google.com/p/azul3d/wrappers/gl/gl21"

const(
    STREAM_DRAW = gl.STREAM_DRAW
    STATIC_DRAW = gl.STATIC_DRAW
    DYNAMIC_DRAW = gl.DYNAMIC_DRAW
)

// Represents a set number of ArrayBuffer
type VertexBuffer struct {
    Buffer gl.Uint
}

func NewVertexBuffer() *VertexBuffer {
    b := VertexBuffer{}

    // Generate buffer number
    gl.GenBuffers(1, &b.Buffer)

    // Create buffer first
    gl.BindBuffer(gl.ARRAY_BUFFER, b.Buffer)

    // Test to ensure they where created
    if gl.IsBuffer(b.Buffer) == gl.FALSE {
        panic("Buffer returned by glGenBuffers() is not a valid buffer, according to glIsBuffer()\nPerhaps you are missing a valid GL context?")
    }

    return &b
}

func (b *VertexBuffer) SetData(data []float32, hint gl.Enum) {
    // Make buffer active
    gl.BindBuffer(gl.ARRAY_BUFFER, b.Buffer)

    // Upload vertex data
    gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(data)), gl.Pointer(&data), hint)
}

func (b *VertexBuffer) Draw(size, vertices int) {
    // Make buffer active
    gl.BindBuffer(gl.ARRAY_BUFFER, b.Buffer)

    //Draw Triangle from VBO - do each time window, view point or data changes
    //Establish its 3 coordinates per vertex with zero stride in this array; necessary here
    gl.VertexPointer(3, gl.FLOAT, 0, nil)

    //Establish array contains vertices (not normals, colours, texture coords etc)
    gl.EnableClientState(gl.VERTEX_ARRAY)

    //Actually draw the triangle, giving the number of vertices provided
    gl.DrawArrays(gl.TRIANGLES, 0, gl.Sizei(vertices))
}

func (b *VertexBuffer) Delete() {
    gl.DeleteBuffers(1, &b.Buffer)
}

/*
//Initialise VBO - do only once, at start of program
//Create a variable to hold the VBO identifier
GLuint triangleVBO;
 
//Vertices of a triangle (counter-clockwise winding)
float data[] = {1.0, 0.0, 1.0, 0.0, 0.0, -1.0, -1.0, 0.0, 1.0};
 
//Create a new VBO and use the variable id to store the VBO id
glGenBuffers(1, &triangleVBO);

 
//Make the new VBO active
glBindBuffer(GL_ARRAY_BUFFER, triangleVBO);
 
//Upload vertex data to the video device
glBufferData(GL_ARRAY_BUFFER, sizeof(data), data, GL_STATIC_DRAW);
 
//Draw Triangle from VBO - do each time window, view point or data changes
//Establish its 3 coordinates per vertex with zero stride in this array; necessary here
glVertexPointer(3, GL_FLOAT, 0, NULL);   
 
//Make the new VBO active. Repeat here incase changed since initialisation
glBindBuffer(GL_ARRAY_BUFFER, triangleVBO);
 
//Establish array contains vertices (not normals, colours, texture coords etc)
glEnableClientState(GL_VERTEX_ARRAY);
 
//Actually draw the triangle, giving the number of vertices provided
glDrawArrays(GL_TRIANGLES, 0, sizeof(data) / sizeof(float) / 3);
 
//Force display to be drawn now
glFlush();
*/
