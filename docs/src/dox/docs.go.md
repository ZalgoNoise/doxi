## /[src](../../src/index.md)/[dox](../dox/index.md)/[docs.go](./docs.go.md)

<pre class="code highlight"><code>
<span id="L1" class="line" lang="go"><a href="#L1">1</a>	</span><span class="comment">package dox</span>
<span id="L2" class="line" lang="go"><a href="#L2">2</a>	</span><span class="comment"></span>
<span id="L3" class="line" lang="go"><a href="#L3">3</a>	</span><span class="comment">import (</span>
<span id="L4" class="line" lang="go"><a href="#L4">4</a>	</span><span class="comment">	"bytes"</span>
<span id="L5" class="line" lang="go"><a href="#L5">5</a>	</span><span class="comment">	"os"</span>
<span id="L6" class="line" lang="go"><a href="#L6">6</a>	</span><span class="comment">	"os/exec"</span>
<span id="L7" class="line" lang="go"><a href="#L7">7</a>	</span><span class="comment">	"strings"</span>
<span id="L8" class="line" lang="go"><a href="#L8">8</a>	</span><span class="comment">)</span>
<span id="L9" class="line" lang="go"><a href="#L9">9</a>	</span><span class="comment"></span>
<span id="L10" class="line" lang="go"><a href="#L10">10</a>	</span><span class="comment">// Docs struct will define a directory for documentation, containing</span>
<span id="L11" class="line" lang="go"><a href="#L11">11</a>	</span><span class="comment">// reference to the actual path to the folder, the project name, and</span>
<span id="L12" class="line" lang="go"><a href="#L12">12</a>	</span><span class="comment">// a list of File objects for each created document</span>
<span id="L13" class="line" lang="go"><a href="#L13">13</a>	</span><span class="comment">type Docs struct {</span>
<span id="L14" class="line" lang="go"><a href="#L14">14</a>	</span><span class="comment">	Path    string `json:"path"`</span>
<span id="L15" class="line" lang="go"><a href="#L15">15</a>	</span><span class="comment">	Proj    string `json:"project"`</span>
<span id="L16" class="line" lang="go"><a href="#L16">16</a>	</span><span class="comment">	Set     string `json:"doc_set"`</span>
<span id="L17" class="line" lang="go"><a href="#L17">17</a>	</span><span class="comment">	Content []File `json:"files"`</span>
<span id="L18" class="line" lang="go"><a href="#L18">18</a>	</span><span class="comment">}</span>
<span id="L19" class="line" lang="go"><a href="#L19">19</a>	</span><span class="comment"></span>
<span id="L20" class="line" lang="go"><a href="#L20">20</a>	</span><span class="comment">// NewDocs function will create a new instance of Docs</span>
<span id="L21" class="line" lang="go"><a href="#L21">21</a>	</span><span class="comment">func NewDocs() *Docs {</span>
<span id="L22" class="line" lang="go"><a href="#L22">22</a>	</span><span class="comment">	d := &Docs{}</span>
<span id="L23" class="line" lang="go"><a href="#L23">23</a>	</span><span class="comment">	return d</span>
<span id="L24" class="line" lang="go"><a href="#L24">24</a>	</span><span class="comment">}</span>
<span id="L25" class="line" lang="go"><a href="#L25">25</a>	</span><span class="comment"></span>
<span id="L26" class="line" lang="go"><a href="#L26">26</a>	</span><span class="comment">// Dir method will scan the Dir object and replay its contents, but inside</span>
<span id="L27" class="line" lang="go"><a href="#L27">27</a>	</span><span class="comment">// the /docs directory. Instead of creating the files straight away, this</span>
<span id="L28" class="line" lang="go"><a href="#L28">28</a>	</span><span class="comment">// part of the process will simply gather the metadata (references for source)</span>
<span id="L29" class="line" lang="go"><a href="#L29">29</a>	</span><span class="comment">// and building the new Files objects (and linking them to the original ones)</span>
<span id="L30" class="line" lang="go"><a href="#L30">30</a>	</span><span class="comment">// in order to branch off of this object with any required actions</span>
<span id="L31" class="line" lang="go"><a href="#L31">31</a>	</span><span class="comment">func (d *Docs) Dir(i *Dir) *Dir {</span>
<span id="L32" class="line" lang="go"><a href="#L32">32</a>	</span><span class="comment"></span>
<span id="L33" class="line" lang="go"><a href="#L33">33</a>	</span><span class="comment">	dir := os.Getenv("PWD")</span>
<span id="L34" class="line" lang="go"><a href="#L34">34</a>	</span><span class="comment"></span>
<span id="L35" class="line" lang="go"><a href="#L35">35</a>	</span><span class="comment">	os.Chdir(d.Path)</span>
<span id="L36" class="line" lang="go"><a href="#L36">36</a>	</span><span class="comment"></span>
<span id="L37" class="line" lang="go"><a href="#L37">37</a>	</span><span class="comment">	defer os.Chdir(dir)</span>
<span id="L38" class="line" lang="go"><a href="#L38">38</a>	</span><span class="comment"></span>
<span id="L39" class="line" lang="go"><a href="#L39">39</a>	</span><span class="comment">	path := d.Path + "/"</span>
<span id="L40" class="line" lang="go"><a href="#L40">40</a>	</span><span class="comment"></span>
<span id="L41" class="line" lang="go"><a href="#L41">41</a>	</span><span class="comment">	for idx, v := range i.Files {</span>
<span id="L42" class="line" lang="go"><a href="#L42">42</a>	</span><span class="comment"></span>
<span id="L43" class="line" lang="go"><a href="#L43">43</a>	</span><span class="comment">		srcRelPath := d.Path + "/src" + v.RelPath.Get()</span>
<span id="L44" class="line" lang="go"><a href="#L44">44</a>	</span><span class="comment">		srcParPath := path + "src/" + v.Parent.Get()</span>
<span id="L45" class="line" lang="go"><a href="#L45">45</a>	</span><span class="comment">		pkgRelPath := d.Path + "/pkg" + v.RelPath.Get()</span>
<span id="L46" class="line" lang="go"><a href="#L46">46</a>	</span><span class="comment">		pkgParPath := path + "pkg/" + v.Parent.Get()</span>
<span id="L47" class="line" lang="go"><a href="#L47">47</a>	</span><span class="comment"></span>
<span id="L48" class="line" lang="go"><a href="#L48">48</a>	</span><span class="comment">		if v.Type == "dir" {</span>
<span id="L49" class="line" lang="go"><a href="#L49">49</a>	</span><span class="comment">			if strings.Contains(srcRelPath, "/docs") {</span>
<span id="L50" class="line" lang="go"><a href="#L50">50</a>	</span><span class="comment">				continue</span>
<span id="L51" class="line" lang="go"><a href="#L51">51</a>	</span><span class="comment">			}</span>
<span id="L52" class="line" lang="go"><a href="#L52">52</a>	</span><span class="comment">			if strings.Contains(srcRelPath, "/bazel-") {</span>
<span id="L53" class="line" lang="go"><a href="#L53">53</a>	</span><span class="comment">				continue</span>
<span id="L54" class="line" lang="go"><a href="#L54">54</a>	</span><span class="comment">			}</span>
<span id="L55" class="line" lang="go"><a href="#L55">55</a>	</span><span class="comment">			if strings.Contains(srcRelPath, "/testdata") {</span>
<span id="L56" class="line" lang="go"><a href="#L56">56</a>	</span><span class="comment">				continue</span>
<span id="L57" class="line" lang="go"><a href="#L57">57</a>	</span><span class="comment">			}</span>
<span id="L58" class="line" lang="go"><a href="#L58">58</a>	</span><span class="comment">			if strings.Contains(srcRelPath, "/.git") {</span>
<span id="L59" class="line" lang="go"><a href="#L59">59</a>	</span><span class="comment">				continue</span>
<span id="L60" class="line" lang="go"><a href="#L60">60</a>	</span><span class="comment">			}</span>
<span id="L61" class="line" lang="go"><a href="#L61">61</a>	</span><span class="comment">			os.Mkdir(srcRelPath, 0755)</span>
<span id="L62" class="line" lang="go"><a href="#L62">62</a>	</span><span class="comment">			os.Mkdir(pkgRelPath, 0755)</span>
<span id="L63" class="line" lang="go"><a href="#L63">63</a>	</span><span class="comment">		}</span>
<span id="L64" class="line" lang="go"><a href="#L64">64</a>	</span><span class="comment"></span>
<span id="L65" class="line" lang="go"><a href="#L65">65</a>	</span><span class="comment">		if v.Type == "file" && v.Ext == ".go" {</span>
<span id="L66" class="line" lang="go"><a href="#L66">66</a>	</span><span class="comment">			if _, err := os.Stat(srcParPath); os.IsNotExist(err) {</span>
<span id="L67" class="line" lang="go"><a href="#L67">67</a>	</span><span class="comment">				os.Mkdir(srcParPath, 0755)</span>
<span id="L68" class="line" lang="go"><a href="#L68">68</a>	</span><span class="comment">				os.Mkdir(pkgParPath, 0755)</span>
<span id="L69" class="line" lang="go"><a href="#L69">69</a>	</span><span class="comment">			}</span>
<span id="L70" class="line" lang="go"><a href="#L70">70</a>	</span><span class="comment">			newSrc, err := os.Create(srcRelPath + ".md")</span>
<span id="L71" class="line" lang="go"><a href="#L71">71</a>	</span><span class="comment">			newPkg, err := os.Create(pkgRelPath + ".md")</span>
<span id="L72" class="line" lang="go"><a href="#L72">72</a>	</span><span class="comment">			defer newSrc.Close()</span>
<span id="L73" class="line" lang="go"><a href="#L73">73</a>	</span><span class="comment">			defer newPkg.Close()</span>
<span id="L74" class="line" lang="go"><a href="#L74">74</a>	</span><span class="comment"></span>
<span id="L75" class="line" lang="go"><a href="#L75">75</a>	</span><span class="comment">			srcF := NewFile()</span>
<span id="L76" class="line" lang="go"><a href="#L76">76</a>	</span><span class="comment">			pkgF := NewFile()</span>
<span id="L77" class="line" lang="go"><a href="#L77">77</a>	</span><span class="comment"></span>
<span id="L78" class="line" lang="go"><a href="#L78">78</a>	</span><span class="comment">			srcF.Init(d.Path, srcRelPath+".md")</span>
<span id="L79" class="line" lang="go"><a href="#L79">79</a>	</span><span class="comment">			i.Files[idx].Link(*srcF)</span>
<span id="L80" class="line" lang="go"><a href="#L80">80</a>	</span><span class="comment">			srcF.Link(v)</span>
<span id="L81" class="line" lang="go"><a href="#L81">81</a>	</span><span class="comment"></span>
<span id="L82" class="line" lang="go"><a href="#L82">82</a>	</span><span class="comment">			pkgF.Init(d.Path, pkgRelPath+".md")</span>
<span id="L83" class="line" lang="go"><a href="#L83">83</a>	</span><span class="comment">			i.Files[idx].Link(*pkgF)</span>
<span id="L84" class="line" lang="go"><a href="#L84">84</a>	</span><span class="comment">			pkgF.Link(v)</span>
<span id="L85" class="line" lang="go"><a href="#L85">85</a>	</span><span class="comment"></span>
<span id="L86" class="line" lang="go"><a href="#L86">86</a>	</span><span class="comment">			d.Content = append(d.Content, *srcF)</span>
<span id="L87" class="line" lang="go"><a href="#L87">87</a>	</span><span class="comment"></span>
<span id="L88" class="line" lang="go"><a href="#L88">88</a>	</span><span class="comment">			if err != nil {</span>
<span id="L89" class="line" lang="go"><a href="#L89">89</a>	</span><span class="comment">				panic(err)</span>
<span id="L90" class="line" lang="go"><a href="#L90">90</a>	</span><span class="comment">			}</span>
<span id="L91" class="line" lang="go"><a href="#L91">91</a>	</span><span class="comment"></span>
<span id="L92" class="line" lang="go"><a href="#L92">92</a>	</span><span class="comment">		}</span>
<span id="L93" class="line" lang="go"><a href="#L93">93</a>	</span><span class="comment"></span>
<span id="L94" class="line" lang="go"><a href="#L94">94</a>	</span><span class="comment">	}</span>
<span id="L95" class="line" lang="go"><a href="#L95">95</a>	</span><span class="comment"></span>
<span id="L96" class="line" lang="go"><a href="#L96">96</a>	</span><span class="comment">	return i</span>
<span id="L97" class="line" lang="go"><a href="#L97">97</a>	</span><span class="comment"></span>
<span id="L98" class="line" lang="go"><a href="#L98">98</a>	</span><span class="comment">}</span>
<span id="L99" class="line" lang="go"><a href="#L99">99</a>	</span><span class="comment"></span>
<span id="L100" class="line" lang="go"><a href="#L100">100</a>	</span><span class="comment">// Source method will create a new SourceCode object for each Go file,</span>
<span id="L101" class="line" lang="go"><a href="#L101">101</a>	</span><span class="comment">// read its contents, get a symlink to the original object, and finally</span>
<span id="L102" class="line" lang="go"><a href="#L102">102</a>	</span><span class="comment">// generate the Markdown file based on this content</span>
<span id="L103" class="line" lang="go"><a href="#L103">103</a>	</span><span class="comment">func (d *Docs) Source(i *Dir) *Dir {</span>
<span id="L104" class="line" lang="go"><a href="#L104">104</a>	</span><span class="comment"></span>
<span id="L105" class="line" lang="go"><a href="#L105">105</a>	</span><span class="comment">	for _, v := range i.Files {</span>
<span id="L106" class="line" lang="go"><a href="#L106">106</a>	</span><span class="comment"></span>
<span id="L107" class="line" lang="go"><a href="#L107">107</a>	</span><span class="comment">		if v.Ext == ".go" {</span>
<span id="L108" class="line" lang="go"><a href="#L108">108</a>	</span><span class="comment">			src := NewSourceCode()</span>
<span id="L109" class="line" lang="go"><a href="#L109">109</a>	</span><span class="comment"></span>
<span id="L110" class="line" lang="go"><a href="#L110">110</a>	</span><span class="comment">			src.Read(v.Path.Get())</span>
<span id="L111" class="line" lang="go"><a href="#L111">111</a>	</span><span class="comment">			src.Link(&v)</span>
<span id="L112" class="line" lang="go"><a href="#L112">112</a>	</span><span class="comment">			src.GenSrc(v.Source)</span>
<span id="L113" class="line" lang="go"><a href="#L113">113</a>	</span><span class="comment">			//src.GenPkg(v.Source)</span>
<span id="L114" class="line" lang="go"><a href="#L114">114</a>	</span><span class="comment">		}</span>
<span id="L115" class="line" lang="go"><a href="#L115">115</a>	</span><span class="comment">	}</span>
<span id="L116" class="line" lang="go"><a href="#L116">116</a>	</span><span class="comment"></span>
<span id="L117" class="line" lang="go"><a href="#L117">117</a>	</span><span class="comment">	return i</span>
<span id="L118" class="line" lang="go"><a href="#L118">118</a>	</span><span class="comment">}</span>
<span id="L119" class="line" lang="go"><a href="#L119">119</a>	</span><span class="comment"></span>
<span id="L120" class="line" lang="go"><a href="#L120">120</a>	</span><span class="comment">func run(args ...string) ([]byte, error) {</span>
<span id="L121" class="line" lang="go"><a href="#L121">121</a>	</span><span class="comment"></span>
<span id="L122" class="line" lang="go"><a href="#L122">122</a>	</span><span class="comment">	cmdPath, err := exec.LookPath(args[0])</span>
<span id="L123" class="line" lang="go"><a href="#L123">123</a>	</span><span class="comment">	if err != nil {</span>
<span id="L124" class="line" lang="go"><a href="#L124">124</a>	</span><span class="comment">		return nil, err</span>
<span id="L125" class="line" lang="go"><a href="#L125">125</a>	</span><span class="comment">	}</span>
<span id="L126" class="line" lang="go"><a href="#L126">126</a>	</span><span class="comment"></span>
<span id="L127" class="line" lang="go"><a href="#L127">127</a>	</span><span class="comment">	cmdFlags := args[1:]</span>
<span id="L128" class="line" lang="go"><a href="#L128">128</a>	</span><span class="comment"></span>
<span id="L129" class="line" lang="go"><a href="#L129">129</a>	</span><span class="comment">	c := exec.Command(cmdPath, cmdFlags...)</span>
<span id="L130" class="line" lang="go"><a href="#L130">130</a>	</span><span class="comment">	var outb, errb bytes.Buffer</span>
<span id="L131" class="line" lang="go"><a href="#L131">131</a>	</span><span class="comment">	c.Stdout = &outb</span>
<span id="L132" class="line" lang="go"><a href="#L132">132</a>	</span><span class="comment">	c.Stderr = &errb</span>
<span id="L133" class="line" lang="go"><a href="#L133">133</a>	</span><span class="comment">	e := c.Run()</span>
<span id="L134" class="line" lang="go"><a href="#L134">134</a>	</span><span class="comment">	if e != nil {</span>
<span id="L135" class="line" lang="go"><a href="#L135">135</a>	</span><span class="comment">		return nil, e</span>
<span id="L136" class="line" lang="go"><a href="#L136">136</a>	</span><span class="comment">	}</span>
<span id="L137" class="line" lang="go"><a href="#L137">137</a>	</span><span class="comment"></span>
<span id="L138" class="line" lang="go"><a href="#L138">138</a>	</span><span class="comment">	return outb.Bytes(), nil</span>
<span id="L139" class="line" lang="go"><a href="#L139">139</a>	</span><span class="comment"></span>
<span id="L140" class="line" lang="go"><a href="#L140">140</a>	</span><span class="comment">}</span>
</code></pre>

_____