## Project Overview

This project provides functionality to securely hash and decrypt auto-incremented IDs using AES encryption. It is designed to be used where sensitive ID information needs to be protected. The primary functions include hashing an ID and decrypting the hashed value back to the original ID, both of which utilize a secret key.

### Key Features

1. **AdjustKey Function**
   - Ensures the secret key is of appropriate length for AES encryption (16, 24, or 32 bytes).
   - If the key is shorter, it pads it; if it's longer, it truncates it.

2. **HashID Function**
   - Takes an auto-incremented ID and a secret key to produce a hashed string.
   - Uses AES encryption in GCM mode for secure and authenticated encryption.
   - The output is a hex-encoded string of the encrypted data.

3. **DecryptID Function**
   - Converts the hex-encoded hash back to the original ID using the same secret key.
   - Verifies the integrity of the encrypted data during decryption.

### Usage

To use these functions, import the package and provide the appropriate inputs:
- `HashID(id int, secret string)`: Generates a hash for the given ID.
- `DecryptID(hash string, secret string)`: Decrypts the hash to retrieve the original ID.

### Example

```go
package main

import (
    "fmt"
    "hash"
)

func main() {
    secret := "your_secret_key"
    id := 12345

    // Hash the ID
    hash, err := hash.HashID(id, secret)
    if err != nil {
        fmt.Println("Error hashing ID:", err)
    } else {
        fmt.Println("Hashed ID:", hash)
    }

    // Decrypt the ID
    originalID, err := hash.DecryptID(hash, secret)
    if err != nil {
        fmt.Println("Error decrypting hash:", err)
    } else {
        fmt.Println("Original ID:", originalID)
    }
}
```

Ensure that the secret key is kept secure, as it is essential for both hashing and decrypting the IDs.


This description provides an overview of the project's purpose, the main functions, and a simple example of how to use them. You can expand or modify it according to your project's needs.
