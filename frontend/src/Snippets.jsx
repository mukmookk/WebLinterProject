class Snippets {
  static java = `public class Main {
    public static void main(String[] args) {
        System.out.println("Hello World");
    }
}`;

  static javascript = `const foo = 1;
let bar = foo;

bar = 9;

console.log(foo, bar);`;

  static python = `import pprint
    
pprint("Hello World" )`;

  static golang = `package main
    
import "fmt"

fmt.Println("Hello World")`;
}

export default Snippets;
