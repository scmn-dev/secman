export function isPlainObject(value: any) {
  if (Object.prototype.toString.call(value) !== "[object Object]") {
    return false;
  }

  const prototype = Object.getPrototypeOf(value);
  return prototype === null || prototype === Object.prototype;
}

export function sortKeys(object: any, options = {}) {
  if (!isPlainObject(object) && !Array.isArray(object)) {
    throw new TypeError("Expected a plain object or array");
  }

  const { deep, compare }: any = options;
  const seenInput: any = [];
  const seenOutput: any = [];

  const deepSortArray = (array: any) => {
    const seenIndex = seenInput.indexOf(array);
    if (seenIndex !== -1) {
      return seenOutput[seenIndex];
    }

    const result: any = [];
    seenInput.push(array);
    seenOutput.push(result);

    result.push(
      ...array.map((item: any) => {
        if (Array.isArray(item)) {
          return deepSortArray(item);
        }

        if (isPlainObject(item)) {
          return _sortKeys(item);
        }

        return item;
      })
    );

    return result;
  };

  const _sortKeys = (object: any) => {
    const seenIndex = seenInput.indexOf(object);
    if (seenIndex !== -1) {
      return seenOutput[seenIndex];
    }

    const result = {};
    const keys = Object.keys(object).sort(compare);

    seenInput.push(object);
    seenOutput.push(result);

    for (const key of keys) {
      const value = object[key];
      let newValue;

      if (deep && Array.isArray(value)) {
        newValue = deepSortArray(value);
      } else {
        newValue = deep && isPlainObject(value) ? _sortKeys(value) : value;
      }

      Object.defineProperty(result, key, {
        ...Object.getOwnPropertyDescriptor(object, key),
        value: newValue,
      });
    }

    return result;
  };

  if (Array.isArray(object)) {
    return deep ? deepSortArray(object) : object.slice();
  }

  return _sortKeys(object);
}
