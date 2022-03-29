export = onetime;
declare function onetime(function_: any, options?: {}): (...arguments_: any[]) => any;
declare namespace onetime {
    function callCount(function_: any): any;
}
