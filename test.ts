let x: string = "123";
console.log(Number(x) + 5); // 123 + 5 = 128

function test(input: string): [string, number] {
  if (input === "error") {
    // 1 adalah error
    return ["", 1];
  }

  // 0 adalah tidak ada error
  return ["hello", 0];
}

const hasil = test("apa saja");
hasil[0]; // hello
hasil[1]; // 0
const [kata, err] = test("apa saja");
kata; // hello
err; // 0
