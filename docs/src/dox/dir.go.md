## /[src](../../src/index.md)/[dox](../dox/index.md)/[dir.go](./dir.go)

<pre class="code highlight"><code>
<span id="L1" class="line" lang="go"><a href="#L1">1</a>	</span><span class="comment">package dox</span>
<span id="L2" class="line" lang="go"><a href="#L2">2</a>	</span><span class="comment"></span>
<span id="L3" class="line" lang="go"><a href="#L3">3</a>	</span><span class="comment">import (</span>
<span id="L4" class="line" lang="go"><a href="#L4">4</a>	</span><span class="comment">	"encoding/json"</span>
<span id="L5" class="line" lang="go"><a href="#L5">5</a>	</span><span class="comment">	"os"</span>
<span id="L6" class="line" lang="go"><a href="#L6">6</a>	</span><span class="comment">	"path/filepath"</span>
<span id="L7" class="line" lang="go"><a href="#L7">7</a>	</span><span class="comment">	"strings"</span>
<span id="L8" class="line" lang="go"><a href="#L8">8</a>	</span><span class="comment">)</span>
<span id="L9" class="line" lang="go"><a href="#L9">9</a>	</span><span class="comment"></span>
<span id="L10" class="line" lang="go"><a href="#L10">10</a>	</span><span class="comment">// Dir struct is the placeholder for the current path which is perceived</span>
<span id="L11" class="line" lang="go"><a href="#L11">11</a>	</span><span class="comment">// as the go project path, and will initiate the process of documenting</span>
<span id="L12" class="line" lang="go"><a href="#L12">12</a>	</span><span class="comment">// Go files</span>
<span id="L13" class="line" lang="go"><a href="#L13">13</a>	</span><span class="comment">type Dir struct {</span>
<span id="L14" class="line" lang="go"><a href="#L14">14</a>	</span><span class="comment">	Path  string</span>
<span id="L15" class="line" lang="go"><a href="#L15">15</a>	</span><span class="comment">	Files []File</span>
<span id="L16" class="line" lang="go"><a href="#L16">16</a>	</span><span class="comment">	Docs  Docs</span>
<span id="L17" class="line" lang="go"><a href="#L17">17</a>	</span><span class="comment">}</span>
<span id="L18" class="line" lang="go"><a href="#L18">18</a>	</span><span class="comment"></span>
<span id="L19" class="line" lang="go"><a href="#L19">19</a>	</span><span class="comment">// New function will create and initialize a new Dir object, and return it</span>
<span id="L20" class="line" lang="go"><a href="#L20">20</a>	</span><span class="comment">func New() *Dir {</span>
<span id="L21" class="line" lang="go"><a href="#L21">21</a>	</span><span class="comment">	d := &Dir{}</span>
<span id="L22" class="line" lang="go"><a href="#L22">22</a>	</span><span class="comment"></span>
<span id="L23" class="line" lang="go"><a href="#L23">23</a>	</span><span class="comment">	d.Init()</span>
<span id="L24" class="line" lang="go"><a href="#L24">24</a>	</span><span class="comment">	return d</span>
<span id="L25" class="line" lang="go"><a href="#L25">25</a>	</span><span class="comment">}</span>
<span id="L26" class="line" lang="go"><a href="#L26">26</a>	</span><span class="comment"></span>
<span id="L27" class="line" lang="go"><a href="#L27">27</a>	</span><span class="comment">// Init method will initialize a new Dir object with the current path</span>
<span id="L28" class="line" lang="go"><a href="#L28">28</a>	</span><span class="comment">// and scan the existing files in the directory, creating File objects</span>
<span id="L29" class="line" lang="go"><a href="#L29">29</a>	</span><span class="comment">// along the way</span>
<span id="L30" class="line" lang="go"><a href="#L30">30</a>	</span><span class="comment">func (d *Dir) Init() *Dir {</span>
<span id="L31" class="line" lang="go"><a href="#L31">31</a>	</span><span class="comment">	d.Path = os.Getenv("PWD")</span>
<span id="L32" class="line" lang="go"><a href="#L32">32</a>	</span><span class="comment">	d.Check("/docs")</span>
<span id="L33" class="line" lang="go"><a href="#L33">33</a>	</span><span class="comment">	d.Check("/docs/src")</span>
<span id="L34" class="line" lang="go"><a href="#L34">34</a>	</span><span class="comment">	d.Check("/docs/pkg")</span>
<span id="L35" class="line" lang="go"><a href="#L35">35</a>	</span><span class="comment"></span>
<span id="L36" class="line" lang="go"><a href="#L36">36</a>	</span><span class="comment">	var files []string</span>
<span id="L37" class="line" lang="go"><a href="#L37">37</a>	</span><span class="comment"></span>
<span id="L38" class="line" lang="go"><a href="#L38">38</a>	</span><span class="comment">	r := d.Path</span>
<span id="L39" class="line" lang="go"><a href="#L39">39</a>	</span><span class="comment">	err := filepath.Walk(</span>
<span id="L40" class="line" lang="go"><a href="#L40">40</a>	</span><span class="comment">		r,</span>
<span id="L41" class="line" lang="go"><a href="#L41">41</a>	</span><span class="comment">		func(p string, i os.FileInfo, err error) error {</span>
<span id="L42" class="line" lang="go"><a href="#L42">42</a>	</span><span class="comment">			files = append(files, p)</span>
<span id="L43" class="line" lang="go"><a href="#L43">43</a>	</span><span class="comment">			return nil</span>
<span id="L44" class="line" lang="go"><a href="#L44">44</a>	</span><span class="comment">		},</span>
<span id="L45" class="line" lang="go"><a href="#L45">45</a>	</span><span class="comment">	)</span>
<span id="L46" class="line" lang="go"><a href="#L46">46</a>	</span><span class="comment">	if err != nil {</span>
<span id="L47" class="line" lang="go"><a href="#L47">47</a>	</span><span class="comment">		panic(err)</span>
<span id="L48" class="line" lang="go"><a href="#L48">48</a>	</span><span class="comment">	}</span>
<span id="L49" class="line" lang="go"><a href="#L49">49</a>	</span><span class="comment">	for _, v := range files {</span>
<span id="L50" class="line" lang="go"><a href="#L50">50</a>	</span><span class="comment">		if v != d.Path {</span>
<span id="L51" class="line" lang="go"><a href="#L51">51</a>	</span><span class="comment">			f := NewFile()</span>
<span id="L52" class="line" lang="go"><a href="#L52">52</a>	</span><span class="comment">			f.Init(d.Path, v)</span>
<span id="L53" class="line" lang="go"><a href="#L53">53</a>	</span><span class="comment">			d.Files = append(d.Files, *f)</span>
<span id="L54" class="line" lang="go"><a href="#L54">54</a>	</span><span class="comment">		}</span>
<span id="L55" class="line" lang="go"><a href="#L55">55</a>	</span><span class="comment">	}</span>
<span id="L56" class="line" lang="go"><a href="#L56">56</a>	</span><span class="comment"></span>
<span id="L57" class="line" lang="go"><a href="#L57">57</a>	</span><span class="comment">	return d</span>
<span id="L58" class="line" lang="go"><a href="#L58">58</a>	</span><span class="comment">}</span>
<span id="L59" class="line" lang="go"><a href="#L59">59</a>	</span><span class="comment"></span>
<span id="L60" class="line" lang="go"><a href="#L60">60</a>	</span><span class="comment">// Check method will look if the /docs directory exist and create it</span>
<span id="L61" class="line" lang="go"><a href="#L61">61</a>	</span><span class="comment">// if it doesn't</span>
<span id="L62" class="line" lang="go"><a href="#L62">62</a>	</span><span class="comment">func (d *Dir) Check(s string) {</span>
<span id="L63" class="line" lang="go"><a href="#L63">63</a>	</span><span class="comment">	path := d.Path + s</span>
<span id="L64" class="line" lang="go"><a href="#L64">64</a>	</span><span class="comment"></span>
<span id="L65" class="line" lang="go"><a href="#L65">65</a>	</span><span class="comment">	if _, err := os.Stat(path); os.IsNotExist(err) {</span>
<span id="L66" class="line" lang="go"><a href="#L66">66</a>	</span><span class="comment"></span>
<span id="L67" class="line" lang="go"><a href="#L67">67</a>	</span><span class="comment">		e := os.Mkdir(path, 0755)</span>
<span id="L68" class="line" lang="go"><a href="#L68">68</a>	</span><span class="comment">		if e != nil {</span>
<span id="L69" class="line" lang="go"><a href="#L69">69</a>	</span><span class="comment">			panic(e)</span>
<span id="L70" class="line" lang="go"><a href="#L70">70</a>	</span><span class="comment">		}</span>
<span id="L71" class="line" lang="go"><a href="#L71">71</a>	</span><span class="comment">	}</span>
<span id="L72" class="line" lang="go"><a href="#L72">72</a>	</span><span class="comment">	if _, err := os.Stat(path); os.IsExist(err) {</span>
<span id="L73" class="line" lang="go"><a href="#L73">73</a>	</span><span class="comment">		os.RemoveAll(path)</span>
<span id="L74" class="line" lang="go"><a href="#L74">74</a>	</span><span class="comment">		e := os.Mkdir(path, 0755)</span>
<span id="L75" class="line" lang="go"><a href="#L75">75</a>	</span><span class="comment">		if e != nil {</span>
<span id="L76" class="line" lang="go"><a href="#L76">76</a>	</span><span class="comment">			panic(e)</span>
<span id="L77" class="line" lang="go"><a href="#L77">77</a>	</span><span class="comment">		}</span>
<span id="L78" class="line" lang="go"><a href="#L78">78</a>	</span><span class="comment">	}</span>
<span id="L79" class="line" lang="go"><a href="#L79">79</a>	</span><span class="comment"></span>
<span id="L80" class="line" lang="go"><a href="#L80">80</a>	</span><span class="comment">}</span>
<span id="L81" class="line" lang="go"><a href="#L81">81</a>	</span><span class="comment"></span>
<span id="L82" class="line" lang="go"><a href="#L82">82</a>	</span><span class="comment">// Run method will execute Doxi for this project, and generate Markdown</span>
<span id="L83" class="line" lang="go"><a href="#L83">83</a>	</span><span class="comment">// documentation for your Go files</span>
<span id="L84" class="line" lang="go"><a href="#L84">84</a>	</span><span class="comment">func (d *Dir) Run() *Dir {</span>
<span id="L85" class="line" lang="go"><a href="#L85">85</a>	</span><span class="comment">	docs := NewDocs()</span>
<span id="L86" class="line" lang="go"><a href="#L86">86</a>	</span><span class="comment">	d.Docs = *docs</span>
<span id="L87" class="line" lang="go"><a href="#L87">87</a>	</span><span class="comment"></span>
<span id="L88" class="line" lang="go"><a href="#L88">88</a>	</span><span class="comment">	path := d.Path + "/docs"</span>
<span id="L89" class="line" lang="go"><a href="#L89">89</a>	</span><span class="comment"></span>
<span id="L90" class="line" lang="go"><a href="#L90">90</a>	</span><span class="comment">	docs.Path = path</span>
<span id="L91" class="line" lang="go"><a href="#L91">91</a>	</span><span class="comment">	proj := strings.Split(d.Path, "/")</span>
<span id="L92" class="line" lang="go"><a href="#L92">92</a>	</span><span class="comment">	docs.Proj = proj[(len(proj) - 1)]</span>
<span id="L93" class="line" lang="go"><a href="#L93">93</a>	</span><span class="comment"></span>
<span id="L94" class="line" lang="go"><a href="#L94">94</a>	</span><span class="comment">	docs.Dir(d)</span>
<span id="L95" class="line" lang="go"><a href="#L95">95</a>	</span><span class="comment"></span>
<span id="L96" class="line" lang="go"><a href="#L96">96</a>	</span><span class="comment">	docs.Source(d)</span>
<span id="L97" class="line" lang="go"><a href="#L97">97</a>	</span><span class="comment"></span>
<span id="L98" class="line" lang="go"><a href="#L98">98</a>	</span><span class="comment">	return d</span>
<span id="L99" class="line" lang="go"><a href="#L99">99</a>	</span><span class="comment"></span>
<span id="L100" class="line" lang="go"><a href="#L100">100</a>	</span><span class="comment">}</span>
<span id="L101" class="line" lang="go"><a href="#L101">101</a>	</span><span class="comment"></span>
<span id="L102" class="line" lang="go"><a href="#L102">102</a>	</span><span class="comment">// Done method will provide output on the execution</span>
<span id="L103" class="line" lang="go"><a href="#L103">103</a>	</span><span class="comment">func (d *Dir) Done() []byte {</span>
<span id="L104" class="line" lang="go"><a href="#L104">104</a>	</span><span class="comment"></span>
<span id="L105" class="line" lang="go"><a href="#L105">105</a>	</span><span class="comment">	json, err := json.Marshal(d.Docs)</span>
<span id="L106" class="line" lang="go"><a href="#L106">106</a>	</span><span class="comment">	if err != nil {</span>
<span id="L107" class="line" lang="go"><a href="#L107">107</a>	</span><span class="comment">		panic(err)</span>
<span id="L108" class="line" lang="go"><a href="#L108">108</a>	</span><span class="comment">	}</span>
<span id="L109" class="line" lang="go"><a href="#L109">109</a>	</span><span class="comment">	return json</span>
<span id="L110" class="line" lang="go"><a href="#L110">110</a>	</span><span class="comment">}</span>
</code></pre>

_____