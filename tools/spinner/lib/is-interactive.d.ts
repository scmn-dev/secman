export function isInteractive({ stream, }?: {
    stream?: NodeJS.WriteStream & {
        fd: 1;
    };
}): boolean;
