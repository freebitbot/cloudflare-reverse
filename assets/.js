function cryptData(color, elem) {
    const characters = "G7Ttuo4rxW6mUa+k2C5eQhN$pqEM8gyFVfcZBSA3dlvRDnY9bPOziwHXJjI-1s0LK";
    const chartCAt = (i) => color.charAt(i);

    let props = {};
    let obj = {};
    let key = "";
    let d = 2;
    let n = 3;
    let p = 2;
    let t = [];
    let data = 0;
    let value = 0;
    let s = 0;

    for (; s < elem.length; s++) {
        const i = elem.charAt(s);

        if (!props.hasOwnProperty(i)) {
            props[i] = n++;
            obj[i] = true;
        }

        const name = key + i;

        if (props.hasOwnProperty(name)) {
            key = name;
        } else {
            let index = ["0", "2", "4", "1", "3"];
            let code = 0;

            while (true) {
                switch (index[code++]) {
                    case "0":
                        if (obj.hasOwnProperty(key)) {
                            if (key.charCodeAt(0) < 256) {
                                for (let o = 0; o < p; o++) {
                                    data = data << 1;
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }

                                let k = key.charCodeAt(0);
                                for (let o = 0; o < 8; o++) {
                                    data = (data << 1) | (k & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = k >> 1;
                                }
                            } else {
                                let dsfsdf = 1;
                                for (let o = 0; o < p; o++) {
                                    data = (data << 1) | dsfsdf;
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    dsfsdf = 0;
                                }

                                let k = key.charCodeAt(0);
                                for (let o = 0; o < 16; o++) {
                                    data = (data << 1) | (dsfsdf & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    dsfsdf = dsfsdf >> 1;
                                }
                            }

                            d--;

                            if (d === 0) {
                                d = Math.pow(2, p);
                                p++;
                            }

                            delete props[key];
                        }
                        break;

                    case "2":
                        if (props.hasOwnProperty(key)) {
                            const k = props[key];
                            for (let o = 0; o < p; o++) {
                                data = (data << 1) | (k & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                                k = k >> 1;
                            }

                            d--;

                            if (d === 0) {
                                d = Math.pow(2, p);
                                p++;
                            }

                            delete props[key];
                        } else {
                            const k = props[key];
                            for (let o = 0; o < p; o++) {
                                data = (data << 1) | (k & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                                k = k >> 1;
                            }

                            d--;

                            if (d === 0) {
                                d = Math.pow(2, p);
                                p++;
                            }

                            delete props[key];
                        }
                        break;

                    case "4":
                        let dfdsfd = 0;
                        for (let o = 0; o < p; o++) {
                            data = (data << 1) | dfdsfd;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                            dfdsfd = 1;
                        }

                        const dfdsf = key.charCodeAt(0);
                        for (let o = 0; o < 16; o++) {
                            data = (data << 1) | (dfdsfd & 1);
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                            dfdsfd = dfdsfd >> 1;
                        }
                        break;

                    case "1":
                        if (obj.hasOwnProperty(key)) {
                            if (key.charCodeAt(0) < 256) {
                                for (let o = 0; o < p; o++) {
                                    data = data << 1;
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }

                                let k = key.charCodeAt(0);
                                for (let o = 0; o < 8; o++) {
                                    data = (data << 1) | (k & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = k >> 1;
                                }
                            } else {
                                let k = 1;
                                for (let o = 0; o < p; o++) {
                                    data = (data << 1) | k;
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = 0;
                                }

                                const edfdsf = key.charCodeAt(0);
                                for (let o = 0; o < 16; o++) {
                                    data = (data << 1) | (k & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = k >> 1;
                                }
                            }
                        }
                        break;

                    default:
                        for (let o = 0; o < p; o++) {
                            data = data << 1;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                        }

                        const k = key.charCodeAt(0);
                        for (let o = 0; o < 8; o++) {
                            data = (data << 1) | (dfdsfd & 1);
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                            dfdsfd = dfdsfd >> 1;
                        }
                        break;
                }

                switch (nodeType) {
                    case "string":
                        if (nodeLength < 32) {
                            for (let i = 8; i > 0; i--) {
                                data = (data << 1) | ((nodeLength >> i) & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            }
                        } else if (nodeLength < 2048) {
                            data = (data << 1) | 0;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }

                            for (let i = 11; i > 0; i--) {
                                data = (data << 1) | ((nodeLength >> i) & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            }
                        } else if (nodeLength < 65536) {
                            data = (data << 1) | 1;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }

                            for (let i = 15; i > 0; i--) {
                                data = (data << 1) | ((nodeLength >> i) & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            }
                        } else {
                            data = (data << 1) | 2;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }

                            for (let i = 23; i > 0; i--) {
                                data = (data << 1) | ((nodeLength >> i) & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            }
                        }

                        for (let i = 0; i < nodeLength; i++) {
                            const c = nodeValue.charCodeAt(i);
                            for (let j = 7; j >= 0; j--) {
                                data = (data << 1) | ((c >> j) & 1);
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            }
                        }
                        break;

                    case "number":
                        if (Math.floor(nodeValue) === nodeValue && isFinite(nodeValue)) {
                            if (nodeValue >= 0 && nodeValue < 2) {
                                data = (data << 1) | nodeValue;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }
                            } else if (nodeValue >= -128 && nodeValue < 128) {
                                data = (data << 1) | 2;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                let k = nodeValue + 128;
                                for (let o = 0; o < 8; o++) {
                                    data = (data << 1) | (k & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = k >> 1;
                                }
                            } else if (nodeValue >= -32768 && nodeValue < 32768) {
                                data = (data << 1) | 3;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                let k = nodeValue + 32768;
                                for (let o = 0; o < 16; o++) {
                                    data = (data << 1) | (k & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                    k = k >> 1;
                                }
                            } else {
                                data = (data << 1) | 4;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                for (let o = 0; o < 32; o++) {
                                    data = (data << 1) | ((nodeValue >> o) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            }
                        } else {
                            data = (data << 1) | 5;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }

                            const strValue = nodeValue.toString();
                            const strLength = strValue.length;

                            if (strLength < 32) {
                                for (let i = 5; i >= 0; i--) {
                                    data = (data << 1) | ((strLength >> i) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            } else if (strLength < 2048) {
                                data = (data << 1) | 0;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                for (let i = 11; i >= 0; i--) {
                                    data = (data << 1) | ((strLength >> i) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            } else if (strLength < 65536) {
                                data = (data << 1) | 1;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                for (let i = 15; i >= 0; i--) {
                                    data = (data << 1) | ((strLength >> i) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            } else {
                                data = (data << 1) | 2;
                                if (value == characters.length - 1) {
                                    value = 0;
                                    t.push(characters.charAt(data));
                                    data = 0;
                                } else {
                                    value++;
                                }

                                for (let i = 23; i >= 0; i--) {
                                    data = (data << 1) | ((strLength >> i) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            }

                            for (let i = 0; i < strLength; i++) {
                                const c = strValue.charCodeAt(i);
                                for (let j = 7; j >= 0; j--) {
                                    data = (data << 1) | ((c >> j) & 1);
                                    if (value == characters.length - 1) {
                                        value = 0;
                                        t.push(characters.charAt(data));
                                        data = 0;
                                    } else {
                                        value++;
                                    }
                                }
                            }
                        }
                        break;

                    case "boolean":
                        if (nodeValue === false) {
                            data = (data << 1) | 0;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                        } else if (nodeValue === true) {
                            data = (data << 1) | 1;
                            if (value == characters.length - 1) {
                                value = 0;
                                t.push(characters.charAt(data));
                                data = 0;
                            } else {
                                value++;
                            }
                        }
                        break;

                    case "undefined":
                        data = (data << 1) | 3;
                        if (value == characters.length - 1) {
                            value = 0;
                            t.push(characters.charAt(data));
                            data = 0;
                        } else {
                            value++;
                        }
                        break;

                    case "null":
                        data = (data << 1) | 4;
                        if (value == characters.length - 1) {
                            value = 0;
                            t.push(characters.charAt(data));
                            data = 0;
                        } else {
                            value++;
                        }
                        break;

                    case "object":
                        data = (data << 1) | 6;
                        if (value == characters.length - 1) {
                            value = 0;
                            t.push(characters.charAt(data));
                            data = 0;
                        } else {
                            value++;
                        }

                        for (const prop in nodeValue) {
                            if (nodeValue.hasOwnProperty(prop)) {
                                stack.push(nodeValue[prop]);
                                stack.push(prop);
                            }
                        }
                        break;
                }
            }

            // Add remaining data to the result
            if (value > 0) {
                data = (data << (8 - value)) | ((1 << (8 - value)) - 1);
                t.push(characters.charAt(data));
            }

            return t.join("");
        }
    };

    if (typeof define === "function" && define.amd) {
        define(function () {
            return jsonpack;
        });
    } else if (typeof module !== "undefined" && module.exports) {
        module.exports = jsonpack;
    } else {
        global.jsonpack = jsonpack;
    }
}