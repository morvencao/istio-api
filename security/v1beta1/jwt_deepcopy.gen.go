// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/jwt.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using JWTRule within kubernetes types, where deepcopy-gen is used.
func (in *JWTRule) DeepCopyInto(out *JWTRule) {
	p := proto.Clone(in).(*JWTRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTRule. Required by controller-gen.
func (in *JWTRule) DeepCopy() *JWTRule {
	if in == nil {
		return nil
	}
	out := new(JWTRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JWTRule. Required by controller-gen.
func (in *JWTRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using JWTHeader within kubernetes types, where deepcopy-gen is used.
func (in *JWTHeader) DeepCopyInto(out *JWTHeader) {
	p := proto.Clone(in).(*JWTHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTHeader. Required by controller-gen.
func (in *JWTHeader) DeepCopy() *JWTHeader {
	if in == nil {
		return nil
	}
	out := new(JWTHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new JWTHeader. Required by controller-gen.
func (in *JWTHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
