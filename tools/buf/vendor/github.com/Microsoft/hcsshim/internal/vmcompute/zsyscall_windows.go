//go:build windows

// Code generated by 'go generate' using "github.com/Microsoft/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package vmcompute

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	return e
}

var (
	modvmcompute = windows.NewLazySystemDLL("vmcompute.dll")

	procHcsCloseComputeSystem              = modvmcompute.NewProc("HcsCloseComputeSystem")
	procHcsCloseProcess                    = modvmcompute.NewProc("HcsCloseProcess")
	procHcsCreateComputeSystem             = modvmcompute.NewProc("HcsCreateComputeSystem")
	procHcsCreateProcess                   = modvmcompute.NewProc("HcsCreateProcess")
	procHcsEnumerateComputeSystems         = modvmcompute.NewProc("HcsEnumerateComputeSystems")
	procHcsGetComputeSystemProperties      = modvmcompute.NewProc("HcsGetComputeSystemProperties")
	procHcsGetProcessInfo                  = modvmcompute.NewProc("HcsGetProcessInfo")
	procHcsGetProcessProperties            = modvmcompute.NewProc("HcsGetProcessProperties")
	procHcsGetServiceProperties            = modvmcompute.NewProc("HcsGetServiceProperties")
	procHcsModifyComputeSystem             = modvmcompute.NewProc("HcsModifyComputeSystem")
	procHcsModifyProcess                   = modvmcompute.NewProc("HcsModifyProcess")
	procHcsModifyServiceSettings           = modvmcompute.NewProc("HcsModifyServiceSettings")
	procHcsOpenComputeSystem               = modvmcompute.NewProc("HcsOpenComputeSystem")
	procHcsOpenProcess                     = modvmcompute.NewProc("HcsOpenProcess")
	procHcsPauseComputeSystem              = modvmcompute.NewProc("HcsPauseComputeSystem")
	procHcsRegisterComputeSystemCallback   = modvmcompute.NewProc("HcsRegisterComputeSystemCallback")
	procHcsRegisterProcessCallback         = modvmcompute.NewProc("HcsRegisterProcessCallback")
	procHcsResumeComputeSystem             = modvmcompute.NewProc("HcsResumeComputeSystem")
	procHcsSaveComputeSystem               = modvmcompute.NewProc("HcsSaveComputeSystem")
	procHcsShutdownComputeSystem           = modvmcompute.NewProc("HcsShutdownComputeSystem")
	procHcsSignalProcess                   = modvmcompute.NewProc("HcsSignalProcess")
	procHcsStartComputeSystem              = modvmcompute.NewProc("HcsStartComputeSystem")
	procHcsTerminateComputeSystem          = modvmcompute.NewProc("HcsTerminateComputeSystem")
	procHcsTerminateProcess                = modvmcompute.NewProc("HcsTerminateProcess")
	procHcsUnregisterComputeSystemCallback = modvmcompute.NewProc("HcsUnregisterComputeSystemCallback")
	procHcsUnregisterProcessCallback       = modvmcompute.NewProc("HcsUnregisterProcessCallback")
)

func hcsCloseComputeSystem(computeSystem HcsSystem) (hr error) {
	hr = procHcsCloseComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsCloseComputeSystem.Addr(), uintptr(computeSystem))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsCloseProcess(process HcsProcess) (hr error) {
	hr = procHcsCloseProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsCloseProcess.Addr(), uintptr(process))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsCreateComputeSystem(id string, configuration string, identity syscall.Handle, computeSystem *HcsSystem, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(configuration)
	if hr != nil {
		return
	}
	return _hcsCreateComputeSystem(_p0, _p1, identity, computeSystem, result)
}

func _hcsCreateComputeSystem(id *uint16, configuration *uint16, identity syscall.Handle, computeSystem *HcsSystem, result **uint16) (hr error) {
	hr = procHcsCreateComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsCreateComputeSystem.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(configuration)), uintptr(identity), uintptr(unsafe.Pointer(computeSystem)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsCreateProcess(computeSystem HcsSystem, processParameters string, processInformation *HcsProcessInformation, process *HcsProcess, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(processParameters)
	if hr != nil {
		return
	}
	return _hcsCreateProcess(computeSystem, _p0, processInformation, process, result)
}

func _hcsCreateProcess(computeSystem HcsSystem, processParameters *uint16, processInformation *HcsProcessInformation, process *HcsProcess, result **uint16) (hr error) {
	hr = procHcsCreateProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsCreateProcess.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(processParameters)), uintptr(unsafe.Pointer(processInformation)), uintptr(unsafe.Pointer(process)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsEnumerateComputeSystems(query string, computeSystems **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcsEnumerateComputeSystems(_p0, computeSystems, result)
}

func _hcsEnumerateComputeSystems(query *uint16, computeSystems **uint16, result **uint16) (hr error) {
	hr = procHcsEnumerateComputeSystems.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsEnumerateComputeSystems.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(computeSystems)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsGetComputeSystemProperties(computeSystem HcsSystem, propertyQuery string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(propertyQuery)
	if hr != nil {
		return
	}
	return _hcsGetComputeSystemProperties(computeSystem, _p0, properties, result)
}

func _hcsGetComputeSystemProperties(computeSystem HcsSystem, propertyQuery *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcsGetComputeSystemProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsGetComputeSystemProperties.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(propertyQuery)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsGetProcessInfo(process HcsProcess, processInformation *HcsProcessInformation, result **uint16) (hr error) {
	hr = procHcsGetProcessInfo.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsGetProcessInfo.Addr(), uintptr(process), uintptr(unsafe.Pointer(processInformation)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsGetProcessProperties(process HcsProcess, processProperties **uint16, result **uint16) (hr error) {
	hr = procHcsGetProcessProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsGetProcessProperties.Addr(), uintptr(process), uintptr(unsafe.Pointer(processProperties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsGetServiceProperties(propertyQuery string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(propertyQuery)
	if hr != nil {
		return
	}
	return _hcsGetServiceProperties(_p0, properties, result)
}

func _hcsGetServiceProperties(propertyQuery *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcsGetServiceProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsGetServiceProperties.Addr(), uintptr(unsafe.Pointer(propertyQuery)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsModifyComputeSystem(computeSystem HcsSystem, configuration string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(configuration)
	if hr != nil {
		return
	}
	return _hcsModifyComputeSystem(computeSystem, _p0, result)
}

func _hcsModifyComputeSystem(computeSystem HcsSystem, configuration *uint16, result **uint16) (hr error) {
	hr = procHcsModifyComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsModifyComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(configuration)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsModifyProcess(process HcsProcess, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcsModifyProcess(process, _p0, result)
}

func _hcsModifyProcess(process HcsProcess, settings *uint16, result **uint16) (hr error) {
	hr = procHcsModifyProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsModifyProcess.Addr(), uintptr(process), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsModifyServiceSettings(settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcsModifyServiceSettings(_p0, result)
}

func _hcsModifyServiceSettings(settings *uint16, result **uint16) (hr error) {
	hr = procHcsModifyServiceSettings.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsModifyServiceSettings.Addr(), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsOpenComputeSystem(id string, computeSystem *HcsSystem, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(id)
	if hr != nil {
		return
	}
	return _hcsOpenComputeSystem(_p0, computeSystem, result)
}

func _hcsOpenComputeSystem(id *uint16, computeSystem *HcsSystem, result **uint16) (hr error) {
	hr = procHcsOpenComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsOpenComputeSystem.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(computeSystem)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsOpenProcess(computeSystem HcsSystem, pid uint32, process *HcsProcess, result **uint16) (hr error) {
	hr = procHcsOpenProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsOpenProcess.Addr(), uintptr(computeSystem), uintptr(pid), uintptr(unsafe.Pointer(process)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsPauseComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsPauseComputeSystem(computeSystem, _p0, result)
}

func _hcsPauseComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsPauseComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsPauseComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsRegisterComputeSystemCallback(computeSystem HcsSystem, callback uintptr, context uintptr, callbackHandle *HcsCallback) (hr error) {
	hr = procHcsRegisterComputeSystemCallback.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsRegisterComputeSystemCallback.Addr(), uintptr(computeSystem), uintptr(callback), uintptr(context), uintptr(unsafe.Pointer(callbackHandle)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsRegisterProcessCallback(process HcsProcess, callback uintptr, context uintptr, callbackHandle *HcsCallback) (hr error) {
	hr = procHcsRegisterProcessCallback.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsRegisterProcessCallback.Addr(), uintptr(process), uintptr(callback), uintptr(context), uintptr(unsafe.Pointer(callbackHandle)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsResumeComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsResumeComputeSystem(computeSystem, _p0, result)
}

func _hcsResumeComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsResumeComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsResumeComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsSaveComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsSaveComputeSystem(computeSystem, _p0, result)
}

func _hcsSaveComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsSaveComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsSaveComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsShutdownComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsShutdownComputeSystem(computeSystem, _p0, result)
}

func _hcsShutdownComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsShutdownComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsShutdownComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsSignalProcess(process HcsProcess, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsSignalProcess(process, _p0, result)
}

func _hcsSignalProcess(process HcsProcess, options *uint16, result **uint16) (hr error) {
	hr = procHcsSignalProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsSignalProcess.Addr(), uintptr(process), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsStartComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsStartComputeSystem(computeSystem, _p0, result)
}

func _hcsStartComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsStartComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsStartComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsTerminateComputeSystem(computeSystem HcsSystem, options string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(options)
	if hr != nil {
		return
	}
	return _hcsTerminateComputeSystem(computeSystem, _p0, result)
}

func _hcsTerminateComputeSystem(computeSystem HcsSystem, options *uint16, result **uint16) (hr error) {
	hr = procHcsTerminateComputeSystem.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsTerminateComputeSystem.Addr(), uintptr(computeSystem), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsTerminateProcess(process HcsProcess, result **uint16) (hr error) {
	hr = procHcsTerminateProcess.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsTerminateProcess.Addr(), uintptr(process), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsUnregisterComputeSystemCallback(callbackHandle HcsCallback) (hr error) {
	hr = procHcsUnregisterComputeSystemCallback.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsUnregisterComputeSystemCallback.Addr(), uintptr(callbackHandle))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcsUnregisterProcessCallback(callbackHandle HcsCallback) (hr error) {
	hr = procHcsUnregisterProcessCallback.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcsUnregisterProcessCallback.Addr(), uintptr(callbackHandle))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}