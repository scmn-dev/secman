# Secman Core

> **Secman Core** is the Backend Infrastructure for Secman products at [**api.secman.dev**](https://api.secman.dev).

**SMCA** server was built with **GoLang**, **Rust**, **TypeScript**, **JavaScript**, **HCL**, **Terraform**, and **Docker**.

## Security

**1. Secman uses The Advanced Encryption Standard (AES) encryption algorithm with Galois/Counter Mode (GCM) symmetric-key cryptographic mode. Passwords encrypted with AES can only be decrypted with the passphrase defined in the _config.yml_ file.**

**2. Endpoints are protected with security middlewares against attacks like XSS.**

**3. Against SQL injection, Secman uses Gorm package to handle database queries which clears all queries.**
