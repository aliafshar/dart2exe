
import 'dart:math';
import 'package:crypto/crypto.dart' as crypto;


String getBytes() {
  final r = new Random();
  final bs = new List<int>.generate(256, (i) => r.nextInt(255));
  return crypto.CryptoUtils.bytesToHex(bs);
}

main() {
  print(
"""

    /      _/_   )      
 __/__. __ / .--_ _., _ 
(_/(_/|/ (<_(__<// /\</_
""");
  print('I am Testing standard lib (math) and external (crypto) dependencies.');
  print('--');
  print(getBytes());
  print('--');
  print('I tested them. Pub works. Goodbye, world!');
}
