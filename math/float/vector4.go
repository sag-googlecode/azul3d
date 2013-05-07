package math

type Vector4 struct {
	X, Y, Z, W Real
}

/*
  643      Removed from Vector4 and made a non-member here because otherwise
  644        OgreMatrix4.h and OgreVector4.h have to try to include and inline each
  645        other, which frankly doesn't work ;)
  646
  647     inline Vector4 operator * (const Vector4& v, const Matrix4& mat)
  648     {
  649         return Vector4(
  650             v.x*mat[0][0] + v.y*mat[1][0] + v.z*mat[2][0] + v.w*mat[3][0],
  651             v.x*mat[0][1] + v.y*mat[1][1] + v.z*mat[2][1] + v.w*mat[3][1],
  652             v.x*mat[0][2] + v.y*mat[1][2] + v.z*mat[2][2] + v.w*mat[3][2],
  653             v.x*mat[0][3] + v.y*mat[1][3] + v.z*mat[2][3] + v.w*mat[3][3]
  654             );
  655     }
  659 }
*/
