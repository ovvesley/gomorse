```
gomorse - A simple CLI tool for Morse code translation.

   ********    *******   ****     ****   *******   *******    ******** ********
  **//////**  **/////** /**/**   **/**  **/////** /**////**  **////// /**///// 
 **      //  **     //**/**//** ** /** **     //**/**   /** /**       /**      
/**         /**      /**/** //***  /**/**      /**/*******  /*********/******* 
/**    *****/**      /**/**  //*   /**/**      /**/**///**  ////////**/**////  
//**  ////**//**     ** /**   /    /**//**     ** /**  //**        /**/**      
 //********  //*******  /**        /** //*******  /**   //** ******** /********
  ////////    ///////   //         //   ///////   //     // ////////  //////// 


Usage: gomorse [options] <input>
Options:
  -m, --morse <input>     Input in Morse code (dots and dashes)
  -l, --literal <input>   Input in literal text
  -h, --help              Show this help message
  -v, --version           Show version info

You can also provide the input directly.
Examples:
  ./gomorse ".... . -.--   .--- ..- -.. ."  (Morse)
Output: HEY JUDE

  ./gomorse "HEY JUDE"                      (Literal)
Output: .... . -.--   .--- ..- -.. .

  ./gomorse -m ".... . -.--   .--- ..- -.. ." (Morse)
Output: HEY JUDE

  ./gomorse -l "HEY JUDE"                      (Literal)
Output: .... . -.--   .--- ..- -.. .
```

### Released

All official releases of this project are available on the Releases page.

There you’ll find versioned builds, changelogs, and downloadable assets for each release. We recommend using the latest stable version for the best experience.

To see all released [click here](https://github.com/ovvesley/gomorse/releases).

### Contribution

The recommended way to contribute to this project is by using a Dev Container. This approach ensures a consistent development environment and simplifies setup—it’s the same method used by projects like Moby and many others.

#### Running with Dev Container

To get started:
1.	Install Visual Studio Code (if you haven’t already).
2.	Install the “Dev Containers” extension from the VS Code marketplace.
3.	Open the project folder in VS Code.
4.	When prompted, click “Reopen in Container”.

VS Code will automatically build the container using the .devcontainer configuration and open a fully functional development environment inside the container.

💡 This setup handles dependencies, tools, and runtime versions out of the box—so you can focus on coding!

For more information or troubleshooting tips, see the official [Dev Containers documentation](https://code.visualstudio.com/docs/devcontainers/containers)

