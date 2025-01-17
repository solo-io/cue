// Code generated by gocode.Generate; DO NOT EDIT.

package filetypes

import (
	"fmt"

	"github.com/solo-io/cue/cue"
	"github.com/solo-io/cue/encoding/gocode/gocodec"
	_ "github.com/solo-io/cue/pkg"
)

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.Value().FieldByName(name, true)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 1713 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xacX\u074f\u0736\x11\x97\xce.P\x11i\u07da\xb7\x02c\x19\b\u0483\xabE\x1e\x8c\x18\v\x1c\f'\xb6\v\xbf4E\x91>\x05\xc1\x81+\x8dv\xd9H\xa4JR\xc9\x1dr\xfb\xd06M\xfbW{\x8b!\xa9o\xdd\xd9\xd7\xfa^nw~\x9c\xe1\xccp>\xf7W\xa7\u007f\x9f\xc5g\xa7\xffD\xf1\xe9\x1fQ\xf4\xf9\xe9\xef\x0f\xe2\xf8#!\x8d\xe52\u01d7\xdcr\xa2\xc7\x0f\xe2\x87\u007fV\xca\xc6gQ\xfc\xf0O\xdc\x1e\u23e2\xf8\x17\xafE\x85&>\xfd\x1cE\xd1oO\xff:\x8b\xe3_\u007f\xf3m\xdebV\x8a*p\xfe\x1c\u0167\x9f\xa2\xe8\xd3\xd3?\x1f\xc4\xf1/\a\xfaOQ|\x16?\xfc#\xaf\x91\x04=tD\x16E\xd1\u06cf3\xd2$\x8e\xcf\xe28\xb1\xd7\r\x9a,o1~\xfb\xf1o\x1a\x9e\u007f\xc7\xf7\b\xbbVT\x05c\x9b\r\xbc\x00\xba\x1fr\xa55\x9aF\xc9\u0080U\xc0\xe1\x0f\xca\x1f\xca\b\xce\xd8c\xfa\xb7\x85\x1fYB\xd7K^\xe3\x16\u009f\xb1Z\xc8=KP\xe6\xaa\x10r\xdf\x03\x8f_\x05\nK\x84\xb4\xa8\x1b\x8d\x96[\xa1\xe4\xf3-<~3\xa1\xb0\xa4T\xba~\u07b3\x12\xf7k\xa5k\x96X\xbe7\xcf\xdd\xc5\xc97\xfe\xa6o\xb7\xfd\x95GvtF\xbc\u0112\xb7\x95\x05a\xc0\x1e\x10HEh\r\x16P*\r\xc6\x16B\x02\x97\x05}R\xad\xcd\xe0\xeb\x03\x82Ak\x85\xdc\x1b(\xb0AY\x90\x14%\a\xeeZ\x15du\x10\xbc\x05g?|2u\xc0y\xfa\xfb\x14n:m\x8e#\u007f\xbe\x91\xa5\x82\x02K!\xd1\xc0A\xfd\x00\u070b\x15\x06\x9c\x9b\xb0p\n\xf5n\xc1\"\xb8\x98\x18\x9d\xb5\xee\x1bK\nn\xf9\xe0\x95s\xab[\x84\x1b(ye\x90%\x1aK\xd4(s4\xdb%\x98_\xe7\x95\aV8\x9dj\x82<O'vJU,Q\r}\xe7\x95g\xf1\xb4\\Ic5\x17\xd2\x0e\xe7\xbeCl\x82_\xcc6\u0404\xccU\xddTh]X\x04Z\xdd(m;\r<\xcdX\x8d\xbc\xee\x94\xf2\xb4B\xe5f0\xd1\u04f8\xb5Z\xecZ\xeb\rp4\xef^z\x17C\x8fG\x0f\xe7up\x8f\\\x88\xd2\xf9\u0082jPso\x89?\x9d\xb1\u0346X\xbf>\xa0A\xb0X7\x15\xb7h\x80kt\x0f \xe95\xac\x82\x1dB+E)\x90\xde\x05\xb8u\xc1\xa0\x95\xb2\xa0J\xb0\aaHH\xaed)\xf6\xad\xbf!c\xee\x02\xf7^B6\xad\xf5q:D\r}\x1b\xe5\xc5y\x9a\xb7H\x11sI\xf4,\xcbX\x92\x1cY\x92Th\xe1\n.\xfc\xf1\xb1;f\xaf\x96L\xfc2\aI\xd2(\x86\xae\xd8p\xb5\t\xaa\xe4-E-\xa5\x9a\xc9L~\xc0\x9a\ae\x88\x17\xaf,J\xe3C\u009dN\xb3\xbf\x1a%\xd3\xf0m\x96\xc3d\ro\xad\xea\xcd9z\x96k^W\xf7e\xb9\x1f\u01d1\xf2>\xc1+\x8a\xaew:\xdcYp\x8b\xc7/?[\xf3y\xf0\xea\xf9\xaa\xcf\xe7\xe0\xdc\u75df\xbd\xc3\xeb\x94\u03c3\u03cf,Qmc\xbb\xc0\xf1Z={\xfa\xe1\xd5z\xf6\xf4\xbez\xe1\xf7T\t\xfe\xf7p\xbe\xfc\xe2\u01477\xe3\x8b\x17\xef0\xa3\x14\x94\xf6c;\n,\xff/3\x9e~\xfe\xe5\xb3\x0f\x9e\x9aN\xea=\xf3\xb3\xebu\xaf\xba4\x85\x9a7\u01b7\x95!u\xa9\x90\x85\xc2\xe8\xa1FSA\xb4\x82\xea\xe0,\xc3\xd3t\xdco/Y\x92\u0498\xd0\x13\xa9\xf3\x12\x81\r\x85`\xa0\x13\xa1\x03\xaa\x80\xf4@EHU\fLSD\u078a\x84\xe21H#\x02\xebK\xc4\n`\xaf\xec\x14\xb0xe\t\u062b\xc1:\a\xec\x15\x91\x1b\xad\xac\x1a\xeb\xeb\bN\x12^\xd9\x0e\xed%M\xd1\xddH\xe7\x01e\t5\x97\xaf^~\xb5\x052\xc4\xe0\u07de8R\x9au\f=\xd3N\xc8f\a\x9b\r\xec\x84\xe4\xfa\xba\xd9\xf5CC7*\x81\x90\x85\xc8}\u007f\xf2\x0fH\xd1\xc0\xadkr\x1a\x1b\x8d\x06%\r.\xc0\xe9i\xf7\x9a\xd7\x19\xeb\a\xad-<\xbaHS/R\xc2t\u0102\x02-\xeaz4\x91\xe4\xa8-\x17\xb2\x93\x03\xe6\xa0\u06aa\xa0>8\x99K6\x1bx\xad4t\xc3\xec\x13p\xb5\xa2\xe6\u05f3\x93\xc0\xa9'\x9b\\\x8b\x9d\xd7\xcfG\xf0\x13\xf8\xe1 \xf2\x03\bk\xb0*]\x0f\xe5\x92Xs%\xbfGm}\xf3\xe5\xf0\xe5_^\x05\x8e\x8c\u0366\xc3~\xe0s3\xe18h\x03\xbdt\xc3\xe9dx\uc1b0\xd9\u0216\x96J\xf9h\xf6#\xa7\xe7J\xfd\xc5ix\x0ez+\x9f]\xb9\xaak\x1a\xd4*!\u0453\xadZ\xe6\x15\x01.\xa3\xbc\x18\x9f\xcc^z/\x99Rx\xafys\x98\xa0\x8e\xe2\xc1\x82\xef'P\xc1\xf7\x1d`\xf9\f\xb1A\xa0\xab\x17?\xb2q5s\xc5\u0301d\xe5\x02\r\xa6\a\xb8Z\xc5+\u007f\x80Rl\x81\xbb\fu\xb0\v\xfe\x05\xee3\xc8\x1d\xe83dqhH5:\u84a5\xd9\xd1l\xecfv\x14\xf6\x80\x9a\x1c\xdd\xe5BH\x17\xe8D<\x015\xc1Y\xd2\xec\xb6p>\xbd\xc5\xff\xa5]\xa6\xa5l9\\\xa4t?\xdc\xc0\x1a\u38cb\xbbY\x1d9X\xb9j`\xda?\x98\xd3cx4/v\xc1\xe3\u0277r\xed\x17n\f\x06\xd26q\x9bqI\x1f\x99IRqw\u035e\x9c\x1ez#\xb1~\x10\xa9\xdd>\x16\xe4\xd2\xc8\xe6\xf1\x05\xbb\x9b\xe6V.\x9cLj!:\xc7\u0674\x104\x1cx\x1fq\xaaA\xc9\x1bq\x8b\xac\x80\xbe\x87 _\x1f\\\x83\xee\u05fb\u0428\xa9@\xf3\xaa\xf2`\x06o,\x14\n\rHeA\u023cj\v\xf4\u06e5\xd25\xbcy\x991w\xce)\xe4v[\xda\xe2/\xfa\x05\xb7\xaf_N{j\u0517k\xd5\x05z-\x83+\xe0\x06R7\x03\xb9O]u\x99\xad]\xf31k\xba\xbc\xcdg\x97\xe9\xaa8G\xa7K\xe3\xa7\x13\xf8w\xf0\u025c\u0092\xd9J9\x977].\xe7\xe8t\xa5\x9c\xa1G\xaa\xf3\xb2\x9bZ\xc7C\xd4\xc2_\xc1G\x8b\xfb\u05ad\x1a\xe4/\n\xf8\xf0\x00\xde\xd7\xe4u*\xdc\xfe\xbf\xcb\xdd\xd9\nO:/|\xbe\xee\xeb;\xb5\x99\xf9q\xdd\u007f\xeb~\x1b\xec\x99\xf4\x1c\x939\x1bF\xb6=\xba\x18B\xa8\xfb9a\xcc<\xeeK\xb4B\xec\xe7~yt\x11\xda\xd8T\xdbN\xad\xc9\xef\x17\xbd]\xe3\xdf-V\rX\xf5K\xaf\u05d1M\xa7\xea\xbeGvI0X0t\xc8a\t\x9ae\x8bO\x12\xb8\xe9\xdem\xbc\x02tz\x8c'\xffA\xf8\xd0>\xa7\u039d\xa8Ai\xe8%O;\xf2\xaa>\xfd\xc1\xa1\u7b1e\x1bt\x18\xb7\x9aw\x1c\xb5\xaa\xbe\xeb\xee\xe1\u0a25\xcfr\xec\xbd\u0180\x89\xf4[f\x82\xd1\v\xccm\xa1F\u007f\x97\x98q\xcf^\x932\xb4\xbc\x99\xf2\vC\x8fl\xda'\xeeQ\xab\xdd2\xe5\x9b\xe0\xf4\x96yW\xbb\u0541w\xf6\xaf\xf7\xe6Zu\xd6<\x9a\x8e,\x8a\xfe\x1b\x00\x00\xff\xff\x8e\xf7,\xe9\xbe\x16\x00\x00")
