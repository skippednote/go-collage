import { useState, useRef } from "react";

const Home = () => {
  const [src, setSrc] = useState();
  const [err, setErr] = useState();
  const [loading, setLoading] = useState();
  const formEl = useRef(null);

  const submit = async (e) => {
    e.preventDefault();
    let [uri, regex, width, gray] = e.currentTarget.elements;
    uri = uri.value;
    regex = regex.value;
    width = width.value;
    gray = gray.checked;

    try {
      setSrc(null);
      setLoading(true);
      const response = await fetch("/api/api", {
        method: "POST",
        body: JSON.stringify({
          uri,
          regex,
          width,
          gray,
        }),
      });
      const collageBlob = await response.blob();
      setSrc(URL.createObjectURL(collageBlob));
      setLoading(false);
    } catch (e) {
      setErr("Failed to load the image");
      console.log(e);
      setLoading(false);
    }
  };

  const axelerant = () => {
    const [url, regex] = Array.from(formEl.current.querySelectorAll("input"));
    url.value = "https://www.axelerant.com/about";
    regex.value = `<div class="emp-avatar">\\s+<img src="(.+jpg)\\?.+" width="300"`;
  };

  return (
    <main className="container">
      <header>
        <h1>GoCollage</h1>
      </header>
      {err && <h4>{err}</h4>}
      <form className="form" onSubmit={submit} ref={formEl}>
        <input name="url" placeholder="url" />
        <input name="regex" placeholder="image regex" />
        <input name="width" placeholder="width" />
        <div className="fields">
          <input
            id="gray"
            name="checkbox"
            placeholder="gray?"
            type="checkbox"
          />
          <label htmlFor="gray"></label>
          <button className="btn" type="submit" disabled={loading}>
            Submit
          </button>
        </div>
      </form>
      <button className="btn try" onClick={axelerant}>
        Try Axelerant Banner?
      </button>
      {src && (
        <>
          <hr />
          <div style={{ textAlign: "center" }}>
            {loading && <h3>Loading...</h3>}
            <a
              href={src}
              style={{ fontSize: "2.5rem" }}
              download="collage.jpg"
              className="btn"
            >
              DOWNLOAD
            </a>
            <br />
            <br />
            <img src={src} />
          </div>
        </>
      )}
    </main>
  );
};
export default Home;
