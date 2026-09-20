import type htmx from "htmx.org";
import type Alpine from "alpinejs";

declare global {
  interface Window {
    htmx: typeof htmx
    Alpine: typeof Alpine
  }
}

export {}