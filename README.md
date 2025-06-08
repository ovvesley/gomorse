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