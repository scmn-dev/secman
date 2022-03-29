declare const _exports: {
    (): void;
    unload: () => void;
    signals: () => any;
    load: () => void;
} | {
    (cb: any, opts: any): () => void;
    unload: () => void;
    signals: () => any;
    load: () => void;
};
export = _exports;
export function signals(): any;
export function unload(): void;
export function load(): void;
