#define UNICODE
#include <windows.h>

#define CINTERFACE
#include <dsound.h>

HRESULT mIDirectSound_Initialize(IDirectSound* i, LPCGUID a) {
	return IDirectSound_Initialize(i, a);
}

HRESULT mIDirectSound_SetCooperativeLevel(IDirectSound* i, HWND a, DWORD b) {
	return IDirectSound_SetCooperativeLevel(i, a, b);
}

void mIDirectSound_AddRef(IDirectSound* i) {
	IDirectSound_AddRef(i);
}

void mIDirectSound_Release(IDirectSound* i) {
	IDirectSound_Release(i);
}

HRESULT mIDirectSound_GetCaps(IDirectSound* i, LPDSCAPS a) {
	return IDirectSound_GetCaps(i, a);
}

HRESULT mIDirectSound_CreateSoundBuffer(IDirectSound* i, LPCDSBUFFERDESC a, LPDIRECTSOUNDBUFFER* b, LPUNKNOWN c) {
	return IDirectSound_CreateSoundBuffer(i, a, b, c);
}



void mIDirectSoundBuffer_AddRef(IDirectSoundBuffer* i) {
	IDirectSoundBuffer_AddRef(i);
}

void mIDirectSoundBuffer_Release(IDirectSoundBuffer* i) {
	IDirectSoundBuffer_Release(i);
}

HRESULT mIDirectSoundBuffer_Lock(IDirectSoundBuffer* i, DWORD a, DWORD b, LPVOID* c, LPDWORD d, LPVOID* e, LPDWORD f, DWORD g) {
	return IDirectSoundBuffer_Lock(i, a, b, c, d, e, f, g);
}

HRESULT mIDirectSoundBuffer_Unlock(IDirectSoundBuffer* i, LPVOID a, DWORD b, LPVOID c, DWORD d) {
	return IDirectSoundBuffer_Unlock(i, a, b, c, d);
}

HRESULT mIDirectSoundBuffer_Play(IDirectSoundBuffer* i, DWORD a, DWORD b, DWORD c) {
	return IDirectSoundBuffer_Play(i, a, b, c);
}

HRESULT mIDirectSoundBuffer_Stop(IDirectSoundBuffer* i) {
	return IDirectSoundBuffer_Stop(i);
}

HRESULT mIDirectSoundBuffer_GetStatus(IDirectSoundBuffer* i, LPDWORD a) {
	return IDirectSoundBuffer_GetStatus(i, a);
}

HRESULT mIDirectSoundBuffer_SetCurrentPosition(IDirectSoundBuffer* i, DWORD a) {
	return IDirectSoundBuffer_SetCurrentPosition(i, a);
}

HRESULT mIDirectSoundBuffer_GetCurrentPosition(IDirectSoundBuffer* i, LPDWORD a, LPDWORD b) {
	return IDirectSoundBuffer_GetCurrentPosition(i, a, b);
}

