(globalThis.TURBOPACK = globalThis.TURBOPACK || []).push([
    typeof document === "object" ? document.currentScript : undefined,
    {},
    {"otherChunks":["static/chunks/[turbopack]_browser_dev_hmr-client_hmr-client_ts_fd44f5a4._.js","static/chunks/node_modules_next_dist_compiled_2ce9398a._.js","static/chunks/node_modules_next_dist_client_8f19e6fb._.js","static/chunks/node_modules_next_dist_2ecbf5fa._.js","static/chunks/node_modules_@swc_helpers_cjs_00636ac3._.js"],"runtimeModuleIds":["[project]/node_modules/next/dist/compiled/@next/react-refresh-utils/dist/runtime.js [app-client] (ecmascript)","[project]/node_modules/next/dist/client/app-next-turbopack.js [app-client] (ecmascript)"]}
]);
(() => {
if (!Array.isArray(globalThis.TURBOPACK)) {
    return;
}

const CHUNK_BASE_PATH = "/_next/";
const CHUNK_SUFFIX_PATH = "";
const RELATIVE_ROOT_PATH = "/ROOT";
const RUNTIME_PUBLIC_PATH = "/_next/";
/**
 * This file contains runtime types and functions that are shared between all
 * TurboPack ECMAScript runtimes.
 *
 * It will be prepended to the runtime code of each runtime.
 */ /* eslint-disable @typescript-eslint/no-unused-vars */ /// <reference path="./runtime-types.d.ts" />
const REEXPORTED_OBJECTS = Symbol("reexported objects");
const hasOwnProperty = Object.prototype.hasOwnProperty;
const toStringTag = typeof Symbol !== "undefined" && Symbol.toStringTag;
function defineProp(obj, name, options) {
    if (!hasOwnProperty.call(obj, name)) Object.defineProperty(obj, name, options);
}
/**
 * Adds the getters to the exports object.
 */ function esm(exports, getters) {
    defineProp(exports, "__esModule", {
        value: true
    });
    if (toStringTag) defineProp(exports, toStringTag, {
        value: "Module"
    });
    for(const key in getters){
        const item = getters[key];
        if (Array.isArray(item)) {
            defineProp(exports, key, {
                get: item[0],
                set: item[1],
                enumerable: true
            });
        } else {
            defineProp(exports, key, {
                get: item,
                enumerable: true
            });
        }
    }
... (truncated for brevity)