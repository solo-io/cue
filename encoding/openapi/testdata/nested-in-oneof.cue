
#A: {
	#B: {}
	#C: {
		b?: #B @protobuf(1,B)
	}
	{} | {
		c: #C @protobuf(1,C)
	}
}