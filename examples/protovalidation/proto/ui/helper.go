package ui

import (
	uiv1 "gitub.com/zodimo/go-compose-examples/gen/ui/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetSelectInputFromEnum(enumDesc protoreflect.EnumDescriptor) *uiv1.SelectInput {
	data := &uiv1.SelectInput{}

	excludeUnspecified := false

	enumOpts := enumDesc.Options()
	if proto.HasExtension(enumOpts, uiv1.E_SelectInput) {
		data.PlaceholderOption = &wrapperspb.StringValue{
			Value: proto.GetExtension(enumOpts, uiv1.E_SelectInput).(*uiv1.SelectInputMetadata).GetPlaceholderOption(),
		}
		data.Label = &wrapperspb.StringValue{
			Value: proto.GetExtension(enumOpts, uiv1.E_SelectInput).(*uiv1.SelectInputMetadata).GetLabel(),
		}
	}

	if proto.HasExtension(enumOpts, uiv1.E_ExcludeUnspecified) {
		excludeUnspecified = proto.GetExtension(enumOpts, uiv1.E_ExcludeUnspecified).(bool)
	}
	values := enumDesc.Values()
	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		opts := v.Options()

		if excludeUnspecified && v.Number() == 0 {
			continue
		}

		label := string(v.Name()) // Fallback to raw name (e.g., "USER_ROLE_UNSPECIFIED")
		if proto.HasExtension(opts, uiv1.E_SelectOptionLabel) {
			maybeLabel, ok := proto.GetExtension(opts, uiv1.E_SelectOptionLabel).(string)
			if ok {
				label = maybeLabel
			}
		}

		data.Options = append(data.Options, &uiv1.SelectInputOption{
			Value: string(v.Name()),
			Label: string(label),
		})
	}
	return data
}
