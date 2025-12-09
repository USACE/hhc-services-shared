import { createSelector } from "redux-bundler";

const base = import.meta.env.BASE_URL;

export default {
    name: "routeUtility",
    doUpdateUrlWithBase: (url) => ({ store }) => {
        const fullUrl = `${base}${url}`;
        store.doUpdateUrl(url.includes(base) ? url : fullUrl);
    },
    selectPathnameMinusBase: createSelector("selectPathname", (pathname) => {
        const parts = pathname.split('/')
        parts.splice(1, 1);
        return parts.join('/')
    })
}
