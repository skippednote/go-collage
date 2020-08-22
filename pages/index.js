const Home = () => {
  const submit = (e) => {
    e.preventDefault();
    let [uri, regex] = e.currentTarget.elements;
    uri = uri.value;
    regex = regex.value.replace(/\\/, "\\\\");

    fetch("http://localhost:8080/", {
      method: "POST",
      body: JSON.stringify({
        uri,
        regex,
      }),
    }).then((r) => r.blob());
  };
  return (
    <main>
      <h1>Home</h1>
      <form onSubmit={submit}>
        <input placeholder="url" />
        <input placeholder="image regex" />
        <button type="submit">Submit</button>
      </form>
    </main>
  );
};
export default Home;
