import cryptojs from "crypto-js";

const ivSize = 128; // 16
const iterations = 100000;
const iv = cryptojs.lib.WordArray.random(ivSize / 8);

export const isBase64 = (value: any) => {
  try {
    return btoa(atob(value)) == value;
  } catch (_) {
    return false;
  }
};

export class CryptoTools {
  static transmissionKey: any;
  static encryptKey: any;

  static sha1(msg: any) {
    return cryptojs.SHA1(msg).toString().toUpperCase();
  }

  static hmac(msg: any, transmissionKey = this.transmissionKey) {
    const encrypted = cryptojs.HmacSHA256(msg, transmissionKey);
    return encrypted.toString();
  }

  static encrypt(message: any, password = this.encryptKey) {
    const salt = cryptojs.lib.WordArray.random(128 / 8);

    const encrypted = cryptojs.AES.encrypt(message, password, {
      iv: iv,
      padding: cryptojs.pad.Pkcs7,
      mode: cryptojs.mode.CBC,
      hasher: cryptojs.algo.SHA256,
    });

    // salt, iv will be hex 32 in length
    // append them to the ciphertext for use  in decryption
    const transitMessage =
      salt.toString() + iv.toString() + encrypted.toString();
    return transitMessage;
  }

  static decrypt(transitMessage: any, password = this.encryptKey) {
    const iv = cryptojs.enc.Hex.parse(transitMessage.substr(32, 32));
    const encrypted = transitMessage.substring(64);

    const decrypted = cryptojs.AES.decrypt(encrypted, password, {
      iv,
      padding: cryptojs.pad.Pkcs7,
      mode: cryptojs.mode.CBC,
      hasher: cryptojs.algo.SHA256,
    });

    return decrypted.toString(cryptojs.enc.Utf8);
  }

  static pbkdf2Encrypt(masterPassword = this.encryptKey, secret: any) {
    const cipher = cryptojs.PBKDF2(masterPassword, secret, {
      hasher: cryptojs.algo.SHA256,
      keySize: 256 / 32,
      iterations,
    });

    return cipher.toString();
  }

  static sha256Encrypt(value: any) {
    return cryptojs.SHA256(value).toString();
  }

  static aesEncrypt(value: any, key = this.transmissionKey) {
    return cryptojs.AES.encrypt(value, key).toString();
  }

  static aesDecrypt(value: any, key = this.transmissionKey) {
    return cryptojs.AES.decrypt(value, key).toString(cryptojs.enc.Utf8);
  }

  static encryptFields(data: any, encryptKey = this.encryptKey) {
    Object.keys(data).forEach((key) => {
      data[key] = this.encrypt(data[key], encryptKey);
    });
  }

  static decryptFields(data: any, keyList: any, encryptKey = this.encryptKey) {
    Object.keys(data).forEach((key) => {
      if (data[key] && keyList.includes(key)) {
        data[key] = this.decrypt(data[key], encryptKey);
      }
    });
  }

  static encryptPayload(
    data: any,
    keyList: any,
    encryptKey = this.encryptKey,
    transmissionKey = this.transmissionKey
  ) {
    Object.keys(data).forEach((key) => {
      if (keyList.includes(key.toLowerCase())) {
        data[key] = this.encrypt(data[key], encryptKey);
      }
    });

    return {
      data: this.aesEncrypt(JSON.stringify(data), transmissionKey),
    };
  }
}
