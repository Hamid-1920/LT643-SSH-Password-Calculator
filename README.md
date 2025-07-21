# 🔐 LT643 SSH Password Generator

![GitHub Release](https://img.shields.io/github/release/Hamid1920/lt643-password-generator?style=flat-square)

Generate SSH passwords for devices (MN6200D, HA6400, TF i60 B1) based on MAC address.


## ⚡ Quick Start

### Prerequisites
- [Go 1.18+](https://go.dev/dl/)

### Installation
```bash
git clone https://github.com/Hamid1920/lt643-password-generator.git
cd LT643-SSH-Password-Calculator
go build
```
## 🚀 Usage
Command-line Mode

### With MAC argument
```bash
./LT643PasswordCalc 00:11:22:33:44:55

# Output:
# LT643 SSH password generator by Hamid1920
# -------------------------------------------------------
# Generated Password for MN6200D: a1b2c3d4e5
# Generated Password for HA6400: d4e5f6a1b2
# Generated Password for TF i60 B1: f6a1b2c3d4
```
### Interactive Mode
```bash
./LT643PasswordCalc
> Enter MAC address: 00-11-22-33-44-55
```
### 🔍 MAC Address Formats
The tool accepts these formats:

001122334455

00:11:22:33:44:55

00-11-22-33-44-55

## ⚙️ Technical Implementation
The algorithm uses PBKDF2 key derivation:

go
password := pbkdf2.Key(
    []byte(uppercaseMAC),
    []byte("Argt3l3comErwin" + deviceType + uppercaseMAC),
    1365,  // Iterations
    6,     // Key length (bytes)
    sha256.New
)

## ⚠️ Important Notes

Valid MAC addresses must be 12 hexadecimal characters

## Device compatibility:

✅ MN6200D

✅ HA6400

✅ TF i60 B1


Disclaimer: Use only on devices you own with proper authorization.
💻 Developed by Hamid1920
🐛 Report issues in GitHub Issues
