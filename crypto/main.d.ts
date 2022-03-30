export declare const isBase64: (value: any) => boolean;
export declare class CryptoTools {
    static transmissionKey: any;
    static encryptKey: any;
    static sha1(msg: any): string;
    static hmac(msg: any, transmissionKey?: any): string;
    static encrypt(message: any, password?: any): string;
    static decrypt(transitMessage: any, password?: any): string;
    static pbkdf2Encrypt(masterPassword: any, secret: any): string;
    static sha256Encrypt(value: any): string;
    static aesEncrypt(value: any, key?: any): string;
    static aesDecrypt(value: any, key?: any): string;
    static encryptFields(data: any, encryptKey?: any): void;
    static decryptFields(data: any, keyList: any, encryptKey?: any): void;
    static encryptPayload(data: any, keyList: any, encryptKey?: any, transmissionKey?: any): {
        data: string;
    };
}
