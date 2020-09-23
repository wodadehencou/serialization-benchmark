
using Go = import "/go.capnp";
@0xc8d62cd5c48a4ab3;
$Go.package("capnp");
$Go.import("github.com/wodadehencou/serialization-benchmark/capnp");

struct Simple {
	strField @0 :Text;
	bytes1Field @1 :Data;
	bytes2Field @2 :Data;
	timestamp @3 :Int64;
}

struct SimpleArr {
	simples @0 :List(Simple);
}
