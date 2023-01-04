package common

func PtrCast[T1, T2 Number](v *T1) *T2 {
	if v == nil {
		return nil
	} else {
		vv := T2(*v)
		return &vv
	}
}

func PtrCastBeforeOpt[T1, T2 Number](v *T1, opt ...func(T1) T1) *T2 {
	if v == nil {
		return nil
	} else {
		vv := *v
		for _, fn := range opt {
			vv = fn(vv)
		}
		vvv := T2(vv)
		return &vvv
	}
}

func PtrCastAfterOpt[T1, T2 Number](v *T1, opt ...func(T2) T2) *T2 {
	if v == nil {
		return nil
	} else {
		vv := T2(*v)
		for _, fn := range opt {
			vv = fn(vv)
		}
		return &vv
	}
}
