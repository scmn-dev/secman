package app

import (
	"io"
	"os"
    "fmt"
	"crypto/aes"
	"crypto/md5"
    "crypto/rand"
    "crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/spf13/cobra"
)

func EncryptCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt a value.",
		Long: "Encrypt a value. with aes, sha256, and md5 algorithms.",
		Run: func(cmd *cobra.Command, args []string) {
			if EncryptOpts.AES {
				if len(EncryptOpts.AESKey) < 32 {
					fmt.Println("AES key must be 32 characters or longer.")
					os.Exit(1)
				} else {
					text := []byte(args[0])
					key := []byte(EncryptOpts.AESKey)

					c, err := aes.NewCipher(key)
					if err != nil {
						fmt.Println(err)
					}

					gcm, err := cipher.NewGCM(c)
					if err != nil {
						fmt.Println(err)
					}

					nonce := make([]byte, gcm.NonceSize())

					if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
						fmt.Println(err)
					}

					fmt.Printf("%x\n", gcm.Seal(nonce, nonce, text, nil))
				}
			} else if EncryptOpts.SHA256 {
				hash := sha256.Sum256([]byte(args[0]))

				fmt.Printf("%x\n", hash)
			} else if EncryptOpts.SHA512 {
				hash := sha512.Sum512([]byte(args[0]))

				fmt.Printf("%x\n", hash)
			} else if EncryptOpts.MD5 {
				hash := md5.Sum([]byte(args[0]))

				fmt.Printf("%x\n", hash)
			} else {
				fmt.Println("No encryption algorithm selected.")
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVarP(&EncryptOpts.AES, "aes", "a", false, "Encrypt with aes algorithm.")
	cmd.Flags().BoolVarP(&EncryptOpts.SHA256, "sha256", "s", false, "Encrypt with sha256 algorithm.")
	cmd.Flags().BoolVarP(&EncryptOpts.SHA512, "sha512", "S", false, "Encrypt with sha512 algorithm.")
	cmd.Flags().BoolVarP(&EncryptOpts.MD5, "md5", "m", false, "Encrypt with md5 algorithm.")
	cmd.Flags().StringVarP(&EncryptOpts.AESKey, "aes-key", "k", "", "Encrypt with aes key.")

	return cmd
}
