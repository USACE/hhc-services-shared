import { SiteWrapper } from "@usace/groundwork";
import "@usace/groundwork/dist/style.css";
import { getNavHelper } from "internal-nav-helper";
import { useConnect } from "redux-bundler-hook";

function App() {
  const { route: Route, doUpdateUrlWithBase } = useConnect("selectRoute", "doUpdateUrlWithBase");
  return (
    <>
      <div
        onClick={getNavHelper((url) => {
          doUpdateUrlWithBase(url);
        })}
      >
        <SiteWrapper>
          <Route />
        </SiteWrapper>
      </div >
    </>
  );
}

export default App;
