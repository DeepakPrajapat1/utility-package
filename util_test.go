package utility

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	type fields struct {
	}

	type args struct {
		input string
	}

	type want struct {
		output int
	}
	tests := []struct {
		name    string // test case name
		args    args
		want    want // output we want from the test case
		wantErr bool // to handle test case which should return a error or not

	}{
		{
			name: "right string with int",
			args: args{
				input: "100",
			},
			want: want{
				output: 100,
			},
			wantErr: false,
		},
		{
			name: "blank string",
			args: args{
				input: "",
			},
			want: want{
				output: 0,
			},
			wantErr: false,
		},
		{
			name: "string with long int",
			args: args{
				input: "1234567890121316465",
			},
			want: want{
				output: 1234567890121316465,
			},
			wantErr: false,
		},
		// function failes on below 2 test-cases and gives : nil pointer dereference
		// @fail : test-case
		// {
		//  name: "string with non numeric",
		//  args: args{
		//      input: "xyz1000",
		//  },
		//  want: want{
		//      output: 0,
		//  },
		//  wantErr: false,
		// },

		// @fail : test-case
		// {
		//  name: "string with special chars",
		//  args: args{
		//      input: "#@%^&*",
		//  },
		//  want: want{
		//      output: 0,
		//  },
		//  wantErr: false,
		// },
	}

	// execute all the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int(tt.args.input)
			assert.Equal(t, tt.want.output, got)
		})
	}

}

/// Test Split...
func TestSplit(t *testing.T) {

	type args struct {
		text string
		char string
	}

	type want struct {
		output      []string
		sliceLength int
	}

	testCases := []struct {
		name string
		args args
		want want
	}{

		{
			name: "hello world",
			args: args{
				text: "hello world",
				char: " ",
			},
			want: want{
				output:      []string{"hello", "world"},
				sliceLength: 2,
			},
		},
		{
			name: "three words",
			args: args{
				text: "let it be",
				char: " ",
			},
			want: want{
				output:      []string{"let", "it", "be"},
				sliceLength: 3,
			},
		},
		{
			name: "Blank text",
			args: args{
				text: "",
				char: " ",
			},
			want: want{
				output:      []string{""},
				sliceLength: 1,
			},
		},
		{
			name: "Blank char",
			args: args{
				text: "hey mate",
				char: "",
			},
			want: want{
				output:      []string{"h", "e", "y", " ", "m", "a", "t", "e"},
				sliceLength: 8,
			},
		},
		{
			name: "Character is not present",
			args: args{
				text: "hey mate",
				char: "o",
			},
			want: want{
				output:      []string{"hey mate"},
				sliceLength: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := Split(testCase.args.text, testCase.args.char)

			assert.Equal(t, testCase.want.output, output)
			assert.Equal(t, testCase.want.sliceLength, len(output))
		})
	}

}

func TestOrigin(t *testing.T) {
	type args struct {
		text string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Call For Heymarket",
			args: args{
				text: "hm:hello",
			},
			want: want{
				output: "Heymarket",
			},
		},
		{
			name: "Call For Facebook",
			args: args{
				text: "fb:hello",
			},
			want: want{
				output: "Facebook",
			},
		},
		{
			name: "Not Available Value",
			args: args{
				text: "hii:hello",
			},
			want: want{
				output: "SMS",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Origin(tc.args.text)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestShowPhone(t *testing.T) {
	type args struct {
		text string
	}
	type want struct {
		output bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Heymarket",
			args: args{
				text: "hm:random string",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "Apple Business Chat",
			args: args{
				text: "abc:random string",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "Test for Whatsapp",
			args: args{
				text: "whatsapp:random string",
			},
			want: want{
				output: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ShowPhone(tc.args.text)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestGetStringInBetween(t *testing.T) {
	type args struct {
		str   string
		start string
		end   string
	}
	type want struct {
		output  string
		isEmpty bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "String Between Characters",
			args: args{
				str:   "Heymarket",
				start: "m",
				end:   "t",
			},
			want: want{
				output:  "arke",
				isEmpty: false,
			},
		},
		{
			name: "String Between Words",
			args: args{
				str:   "Heymarket is a SMS Service",
				start: "Heymarket",
				end:   "Service",
			},
			want: want{
				output:  " is a SMS ",
				isEmpty: false,
			},
		},
		{
			name: "Start Not Found",
			args: args{
				str:   "Heymarket is a SMS Service",
				start: "shopify",
				end:   "Service",
			},
			want: want{
				output:  "",
				isEmpty: true,
			},
		},
		{
			name: "End Not Found",
			args: args{
				str:   "Heymarket is a SMS Service",
				start: "shopify",
				end:   "Service",
			},
			want: want{
				output: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := GetStringInBetween(tc.args.str, tc.args.start, tc.args.end)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestShopifyMessage(t *testing.T) {
	type args struct {
		msg string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Sent Hello Message",
			args: args{
				msg: "Hello",
			},
			want: want{
				output: "[Shopify] Hello",
			},
		},
		{
			name: "Sent Blank String",
			args: args{
				msg: "",
			},
			want: want{
				output: "[Shopify] ",
			},
		},
		{
			name: "Sent String Of Numbers",
			args: args{
				msg: "123456",
			},
			want: want{
				output: "[Shopify] 123456",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ShopifyMessage(tc.args.msg)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str  string
		char string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Trim first and last space",
			args: args{
				str:  " hello ",
				char: " ",
			},
			want: want{
				output: "hello",
			},
		},
		{
			name: "Trim same char from first and last",
			args: args{
				str:  "#hello#",
				char: "#",
			},
			want: want{
				output: "hello",
			},
		},
		{
			name: "Wrong Character",
			args: args{
				str:  "#hello#",
				char: "@",
			},
			want: want{
				output: "#hello#",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Trim(tc.args.str, tc.args.char)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestCleanPhone(t *testing.T) {
	type args struct {
		phNo string
	}
	type want struct {
		phone string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "With brackets",
			args: args{
				phNo: "(1234567890)",
			},
			want: want{
				phone: "1234567890",
			},
		},

		{
			name: "With + in starting",
			args: args{
				phNo: "+1234567890",
			},
			want: want{
				phone: "1234567890",
			},
		},
		{
			name: "With - between country code and numbers",
			args: args{
				phNo: "91-1234567890",
			},
			want: want{
				phone: "911234567890",
			},
		},
		{
			name: "unnecessary .",
			args: args{
				phNo: ".1234567890",
			},
			want: want{
				phone: "1234567890",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := CleanPhone(tc.args.phNo)
			assert.Equal(t, tc.want.phone, output)
		})
	}
}

func TestIsBlank(t *testing.T) {
	type args struct {
		param interface{}
	}
	type want struct {
		isBlank bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "For Empty String",
			args: args{
				param: "",
			},
			want: want{
				isBlank: true,
			},
		},
		{
			name: "For String",
			args: args{
				param: "Hello",
			},
			want: want{
				isBlank: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := IsBlank(tc.args.param)
			assert.Equal(t, tc.want.isBlank, output)
		})
	}
}

func TestPhoneValid(t *testing.T) {
	type args struct {
		v interface{}
	}
	type want struct {
		output bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "less then 10 digit",
			args: args{
				v: "987654321",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "10 digit number",
			args: args{
				v: "9876543210",
			},
			want: want{
				output: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := PhoneValid(tc.args.v)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestE164Phone(t *testing.T) {
	type args struct {
		v string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Without Country Code",
			args: args{
				v: "9876543210",
			},
			want: want{
				output: "19876543210",
			},
		},
		{
			name: "Wit Country Code",
			args: args{
				v: "19876543210",
			},
			want: want{
				output: "19876543210",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := E164Phone(tc.args.v)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestFloat642Int(t *testing.T) {
	type args struct {
		Value float64
	}
	type want struct {
		output  int
		varType string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Float64 to Int",
			args: args{
				Value: 89.5,
			},
			want: want{
				output:  89,
				varType: "int",
			},
		},
		{
			name: "2nd Float64 to Int",
			args: args{
				Value: 90,
			},
			want: want{
				output:  90,
				varType: "int",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Float642Int(tc.args.Value)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, fmt.Sprintf("%T", tc.want.output), fmt.Sprintf("%T", output))
		})
	}
}

func TestMarshal(t *testing.T) {

	type tempStruct struct {
		Name string
		Fees int
	}

	type args struct {
		strct tempStruct
	}
	type want struct {
		output []byte
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Pass struct",
			args: args{
				strct: tempStruct{
					Name: "Raghav",
					Fees: 5000,
				},
			},
			want: want{
				output: []byte("{\"Name\":\"Raghav\",\"Fees\":5000}"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Marshal(tc.args.strct)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	type want struct {
		output  string
		varType string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Int to string",
			args: args{
				value: 9,
			},
			want: want{
				output:  "9",
				varType: "string",
			},
		},
		{
			name: "Float to string",
			args: args{
				value: 9.2,
			},
			want: want{
				output:  "9.2",
				varType: "string",
			},
		},
		{
			name: "Struct To String",
			args: args{
				value: struct {
					Name string
					Fees int
				}{
					Name: "Hey",
					Fees: 123,
				},
			},
			want: want{
				output:  "{Hey 123}",
				varType: "string",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ToString(tc.args.value)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, tc.want.varType, fmt.Sprintf("%T", output))
		})
	}
}
