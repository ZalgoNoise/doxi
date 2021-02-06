## /[src](../../src/index.md)/[dox](../dox/index.md)/[data.go](./data.go.md)

<pre class="code highlight"><code>
<span id="L1" class="line" lang="go"><a href="#L1">1</a>	</span><span class="comment">package dox</span>
<span id="L2" class="line" lang="go"><a href="#L2">2</a>	</span><span class="comment"></span>
<span id="L3" class="line" lang="go"><a href="#L3">3</a>	</span><span class="comment">import (</span>
<span id="L4" class="line" lang="go"><a href="#L4">4</a>	</span><span class="comment">	"bufio"</span>
<span id="L5" class="line" lang="go"><a href="#L5">5</a>	</span><span class="comment">	"bytes"</span>
<span id="L6" class="line" lang="go"><a href="#L6">6</a>	</span><span class="comment">	"io/ioutil"</span>
<span id="L7" class="line" lang="go"><a href="#L7">7</a>	</span><span class="comment">	"os"</span>
<span id="L8" class="line" lang="go"><a href="#L8">8</a>	</span><span class="comment">	"strconv"</span>
<span id="L9" class="line" lang="go"><a href="#L9">9</a>	</span><span class="comment">	"strings"</span>
<span id="L10" class="line" lang="go"><a href="#L10">10</a>	</span><span class="comment">)</span>
<span id="L11" class="line" lang="go"><a href="#L11">11</a>	</span><span class="comment"></span>
<span id="L12" class="line" lang="go"><a href="#L12">12</a>	</span><span class="comment">var newLine = []byte(`</span>
<span id="L13" class="line" lang="go"><a href="#L13">13</a>	</span><span class="comment">`)</span>
<span id="L14" class="line" lang="go"><a href="#L14">14</a>	</span><span class="comment"></span>
<span id="L15" class="line" lang="go"><a href="#L15">15</a>	</span><span class="comment">// SourceCode struct will contain one instance of a source code file</span>
<span id="L16" class="line" lang="go"><a href="#L16">16</a>	</span><span class="comment">// with all required parameters to build one, plus a pointer to the</span>
<span id="L17" class="line" lang="go"><a href="#L17">17</a>	</span><span class="comment">// original file</span>
<span id="L18" class="line" lang="go"><a href="#L18">18</a>	</span><span class="comment">type SourceCode struct {</span>
<span id="L19" class="line" lang="go"><a href="#L19">19</a>	</span><span class="comment">	Content []string</span>
<span id="L20" class="line" lang="go"><a href="#L20">20</a>	</span><span class="comment">	Buffer  []byte</span>
<span id="L21" class="line" lang="go"><a href="#L21">21</a>	</span><span class="comment">	Lines   int</span>
<span id="L22" class="line" lang="go"><a href="#L22">22</a>	</span><span class="comment">	Proj    string</span>
<span id="L23" class="line" lang="go"><a href="#L23">23</a>	</span><span class="comment">	File    string</span>
<span id="L24" class="line" lang="go"><a href="#L24">24</a>	</span><span class="comment">	Path    string</span>
<span id="L25" class="line" lang="go"><a href="#L25">25</a>	</span><span class="comment">	RelPath string</span>
<span id="L26" class="line" lang="go"><a href="#L26">26</a>	</span><span class="comment">	Ref     *File</span>
<span id="L27" class="line" lang="go"><a href="#L27">27</a>	</span><span class="comment">}</span>
<span id="L28" class="line" lang="go"><a href="#L28">28</a>	</span><span class="comment"></span>
<span id="L29" class="line" lang="go"><a href="#L29">29</a>	</span><span class="comment">// NewSourceCode function will create a new instance of SourceCode</span>
<span id="L30" class="line" lang="go"><a href="#L30">30</a>	</span><span class="comment">func NewSourceCode() *SourceCode {</span>
<span id="L31" class="line" lang="go"><a href="#L31">31</a>	</span><span class="comment">	return &SourceCode{}</span>
<span id="L32" class="line" lang="go"><a href="#L32">32</a>	</span><span class="comment">}</span>
<span id="L33" class="line" lang="go"><a href="#L33">33</a>	</span><span class="comment"></span>
<span id="L34" class="line" lang="go"><a href="#L34">34</a>	</span><span class="comment">// Read method will define the SourceCode object as per the file in the</span>
<span id="L35" class="line" lang="go"><a href="#L35">35</a>	</span><span class="comment">// provided path</span>
<span id="L36" class="line" lang="go"><a href="#L36">36</a>	</span><span class="comment">func (s *SourceCode) Read(path string) *SourceCode {</span>
<span id="L37" class="line" lang="go"><a href="#L37">37</a>	</span><span class="comment">	var content []string</span>
<span id="L38" class="line" lang="go"><a href="#L38">38</a>	</span><span class="comment"></span>
<span id="L39" class="line" lang="go"><a href="#L39">39</a>	</span><span class="comment">	file, err := os.Open(path)</span>
<span id="L40" class="line" lang="go"><a href="#L40">40</a>	</span><span class="comment">	if err != nil {</span>
<span id="L41" class="line" lang="go"><a href="#L41">41</a>	</span><span class="comment">		panic(err)</span>
<span id="L42" class="line" lang="go"><a href="#L42">42</a>	</span><span class="comment">	}</span>
<span id="L43" class="line" lang="go"><a href="#L43">43</a>	</span><span class="comment">	defer file.Close()</span>
<span id="L44" class="line" lang="go"><a href="#L44">44</a>	</span><span class="comment"></span>
<span id="L45" class="line" lang="go"><a href="#L45">45</a>	</span><span class="comment">	scanner := bufio.NewScanner(file)</span>
<span id="L46" class="line" lang="go"><a href="#L46">46</a>	</span><span class="comment">	for scanner.Scan() {</span>
<span id="L47" class="line" lang="go"><a href="#L47">47</a>	</span><span class="comment">		content = append(content, scanner.Text())</span>
<span id="L48" class="line" lang="go"><a href="#L48">48</a>	</span><span class="comment">	}</span>
<span id="L49" class="line" lang="go"><a href="#L49">49</a>	</span><span class="comment"></span>
<span id="L50" class="line" lang="go"><a href="#L50">50</a>	</span><span class="comment">	if err := scanner.Err(); err != nil {</span>
<span id="L51" class="line" lang="go"><a href="#L51">51</a>	</span><span class="comment">		panic(err)</span>
<span id="L52" class="line" lang="go"><a href="#L52">52</a>	</span><span class="comment">	}</span>
<span id="L53" class="line" lang="go"><a href="#L53">53</a>	</span><span class="comment"></span>
<span id="L54" class="line" lang="go"><a href="#L54">54</a>	</span><span class="comment">	s.Content = content</span>
<span id="L55" class="line" lang="go"><a href="#L55">55</a>	</span><span class="comment">	s.Lines = len(content)</span>
<span id="L56" class="line" lang="go"><a href="#L56">56</a>	</span><span class="comment"></span>
<span id="L57" class="line" lang="go"><a href="#L57">57</a>	</span><span class="comment">	return s</span>
<span id="L58" class="line" lang="go"><a href="#L58">58</a>	</span><span class="comment">}</span>
<span id="L59" class="line" lang="go"><a href="#L59">59</a>	</span><span class="comment"></span>
<span id="L60" class="line" lang="go"><a href="#L60">60</a>	</span><span class="comment">// Link method will point a SourceCode object to its original File</span>
<span id="L61" class="line" lang="go"><a href="#L61">61</a>	</span><span class="comment">func (s *SourceCode) Link(f *File) *SourceCode {</span>
<span id="L62" class="line" lang="go"><a href="#L62">62</a>	</span><span class="comment">	s.Proj = f.Proj.Get()</span>
<span id="L63" class="line" lang="go"><a href="#L63">63</a>	</span><span class="comment">	s.Ref = f</span>
<span id="L64" class="line" lang="go"><a href="#L64">64</a>	</span><span class="comment">	s.RelPath = f.RelPath.Get()</span>
<span id="L65" class="line" lang="go"><a href="#L65">65</a>	</span><span class="comment">	return s</span>
<span id="L66" class="line" lang="go"><a href="#L66">66</a>	</span><span class="comment">}</span>
<span id="L67" class="line" lang="go"><a href="#L67">67</a>	</span><span class="comment"></span>
<span id="L68" class="line" lang="go"><a href="#L68">68</a>	</span><span class="comment">// SetPath method will point the files to the set /docs/* folder</span>
<span id="L69" class="line" lang="go"><a href="#L69">69</a>	</span><span class="comment">func (s *SourceCode) SetPath(p string) string {</span>
<span id="L70" class="line" lang="go"><a href="#L70">70</a>	</span><span class="comment">	paths := strings.Split(s.Ref.Path.Get(), s.Ref.RelPath.Get())</span>
<span id="L71" class="line" lang="go"><a href="#L71">71</a>	</span><span class="comment">	srcPath := paths[0] + "/docs/" + p + s.Ref.RelPath.Get() + ".md"</span>
<span id="L72" class="line" lang="go"><a href="#L72">72</a>	</span><span class="comment"></span>
<span id="L73" class="line" lang="go"><a href="#L73">73</a>	</span><span class="comment">	return srcPath</span>
<span id="L74" class="line" lang="go"><a href="#L74">74</a>	</span><span class="comment">}</span>
<span id="L75" class="line" lang="go"><a href="#L75">75</a>	</span><span class="comment"></span>
<span id="L76" class="line" lang="go"><a href="#L76">76</a>	</span><span class="comment">// GenSrc method will create the /docs/src content for the SourceCode</span>
<span id="L77" class="line" lang="go"><a href="#L77">77</a>	</span><span class="comment">// generated document</span>
<span id="L78" class="line" lang="go"><a href="#L78">78</a>	</span><span class="comment">func (s *SourceCode) GenSrc(f *File) *SourceCode {</span>
<span id="L79" class="line" lang="go"><a href="#L79">79</a>	</span><span class="comment">	s.File = f.Name.Get()</span>
<span id="L80" class="line" lang="go"><a href="#L80">80</a>	</span><span class="comment">	s.Path = s.SetPath("src")</span>
<span id="L81" class="line" lang="go"><a href="#L81">81</a>	</span><span class="comment"></span>
<span id="L82" class="line" lang="go"><a href="#L82">82</a>	</span><span class="comment">	header := s.GenHeader("src")</span>
<span id="L83" class="line" lang="go"><a href="#L83">83</a>	</span><span class="comment"></span>
<span id="L84" class="line" lang="go"><a href="#L84">84</a>	</span><span class="comment">	code := s.GenCode()</span>
<span id="L85" class="line" lang="go"><a href="#L85">85</a>	</span><span class="comment"></span>
<span id="L86" class="line" lang="go"><a href="#L86">86</a>	</span><span class="comment">	var buf []byte</span>
<span id="L87" class="line" lang="go"><a href="#L87">87</a>	</span><span class="comment">	s.Buffer = byteJoin(buf, header, code)</span>
<span id="L88" class="line" lang="go"><a href="#L88">88</a>	</span><span class="comment"></span>
<span id="L89" class="line" lang="go"><a href="#L89">89</a>	</span><span class="comment">	err := ioutil.WriteFile(s.Path, s.Buffer, 0644)</span>
<span id="L90" class="line" lang="go"><a href="#L90">90</a>	</span><span class="comment">	if err != nil {</span>
<span id="L91" class="line" lang="go"><a href="#L91">91</a>	</span><span class="comment">		panic(err)</span>
<span id="L92" class="line" lang="go"><a href="#L92">92</a>	</span><span class="comment">	}</span>
<span id="L93" class="line" lang="go"><a href="#L93">93</a>	</span><span class="comment"></span>
<span id="L94" class="line" lang="go"><a href="#L94">94</a>	</span><span class="comment">	return s</span>
<span id="L95" class="line" lang="go"><a href="#L95">95</a>	</span><span class="comment">}</span>
<span id="L96" class="line" lang="go"><a href="#L96">96</a>	</span><span class="comment"></span>
<span id="L97" class="line" lang="go"><a href="#L97">97</a>	</span><span class="comment">// GenPkg method will create the /docs/pkg content for the SourceCode</span>
<span id="L98" class="line" lang="go"><a href="#L98">98</a>	</span><span class="comment">// generated document</span>
<span id="L99" class="line" lang="go"><a href="#L99">99</a>	</span><span class="comment">func (s *SourceCode) GenPkg(f *File) *SourceCode {</span>
<span id="L100" class="line" lang="go"><a href="#L100">100</a>	</span><span class="comment">	s.File = f.Name.Get()</span>
<span id="L101" class="line" lang="go"><a href="#L101">101</a>	</span><span class="comment">	s.Path = s.SetPath("pkg")</span>
<span id="L102" class="line" lang="go"><a href="#L102">102</a>	</span><span class="comment"></span>
<span id="L103" class="line" lang="go"><a href="#L103">103</a>	</span><span class="comment">	//header := s.GenHeader("pkg")</span>
<span id="L104" class="line" lang="go"><a href="#L104">104</a>	</span><span class="comment"></span>
<span id="L105" class="line" lang="go"><a href="#L105">105</a>	</span><span class="comment">	return s</span>
<span id="L106" class="line" lang="go"><a href="#L106">106</a>	</span><span class="comment">}</span>
<span id="L107" class="line" lang="go"><a href="#L107">107</a>	</span><span class="comment"></span>
<span id="L108" class="line" lang="go"><a href="#L108">108</a>	</span><span class="comment">func (s *SourceCode) GenHeader(docs string) []byte {</span>
<span id="L109" class="line" lang="go"><a href="#L109">109</a>	</span><span class="comment"></span>
<span id="L110" class="line" lang="go"><a href="#L110">110</a>	</span><span class="comment">	hPath := strings.Split(s.RelPath, "/")</span>
<span id="L111" class="line" lang="go"><a href="#L111">111</a>	</span><span class="comment">	hPath[0] = docs</span>
<span id="L112" class="line" lang="go"><a href="#L112">112</a>	</span><span class="comment"></span>
<span id="L113" class="line" lang="go"><a href="#L113">113</a>	</span><span class="comment">	var headerURL []string</span>
<span id="L114" class="line" lang="go"><a href="#L114">114</a>	</span><span class="comment"></span>
<span id="L115" class="line" lang="go"><a href="#L115">115</a>	</span><span class="comment">	for i := 0; i < len(hPath); i++ {</span>
<span id="L116" class="line" lang="go"><a href="#L116">116</a>	</span><span class="comment">		headerURL = append(headerURL, mdRelURL(hPath[i], (len(hPath)-1-i)))</span>
<span id="L117" class="line" lang="go"><a href="#L117">117</a>	</span><span class="comment">	}</span>
<span id="L118" class="line" lang="go"><a href="#L118">118</a>	</span><span class="comment"></span>
<span id="L119" class="line" lang="go"><a href="#L119">119</a>	</span><span class="comment">	var hURL string</span>
<span id="L120" class="line" lang="go"><a href="#L120">120</a>	</span><span class="comment"></span>
<span id="L121" class="line" lang="go"><a href="#L121">121</a>	</span><span class="comment">	for _, v := range headerURL {</span>
<span id="L122" class="line" lang="go"><a href="#L122">122</a>	</span><span class="comment">		hURL += "/" + v</span>
<span id="L123" class="line" lang="go"><a href="#L123">123</a>	</span><span class="comment">	}</span>
<span id="L124" class="line" lang="go"><a href="#L124">124</a>	</span><span class="comment"></span>
<span id="L125" class="line" lang="go"><a href="#L125">125</a>	</span><span class="comment">	return []byte("## " + hURL + "\n\n")</span>
<span id="L126" class="line" lang="go"><a href="#L126">126</a>	</span><span class="comment"></span>
<span id="L127" class="line" lang="go"><a href="#L127">127</a>	</span><span class="comment">}</span>
<span id="L128" class="line" lang="go"><a href="#L128">128</a>	</span><span class="comment"></span>
<span id="L129" class="line" lang="go"><a href="#L129">129</a>	</span><span class="comment">func (s *SourceCode) GenCode() []byte {</span>
<span id="L130" class="line" lang="go"><a href="#L130">130</a>	</span><span class="comment">	var gen []string</span>
<span id="L131" class="line" lang="go"><a href="#L131">131</a>	</span><span class="comment">	var buf []byte</span>
<span id="L132" class="line" lang="go"><a href="#L132">132</a>	</span><span class="comment"></span>
<span id="L133" class="line" lang="go"><a href="#L133">133</a>	</span><span class="comment">	codeFmt := []byte("<pre class=" + `"code highlight"` + "><code>\n")</span>
<span id="L134" class="line" lang="go"><a href="#L134">134</a>	</span><span class="comment"></span>
<span id="L135" class="line" lang="go"><a href="#L135">135</a>	</span><span class="comment">	buf = byteJoin(buf, codeFmt)</span>
<span id="L136" class="line" lang="go"><a href="#L136">136</a>	</span><span class="comment">	for idx, v := range s.Content {</span>
<span id="L137" class="line" lang="go"><a href="#L137">137</a>	</span><span class="comment">		line := strconv.Itoa(idx + 1)</span>
<span id="L138" class="line" lang="go"><a href="#L138">138</a>	</span><span class="comment">		entry := `<span id="L` + line + `" class="line" lang="go"><a href="#L` + line + `">` + line + `</a>	</span><span class="comment">` + v + `</span>`</span>
<span id="L139" class="line" lang="go"><a href="#L139">139</a>	</span><span class="comment"></span>
<span id="L140" class="line" lang="go"><a href="#L140">140</a>	</span><span class="comment">		gen = append(gen, entry)</span>
<span id="L141" class="line" lang="go"><a href="#L141">141</a>	</span><span class="comment">		buf = byteJoin(buf, []byte(entry), newLine)</span>
<span id="L142" class="line" lang="go"><a href="#L142">142</a>	</span><span class="comment">	}</span>
<span id="L143" class="line" lang="go"><a href="#L143">143</a>	</span><span class="comment"></span>
<span id="L144" class="line" lang="go"><a href="#L144">144</a>	</span><span class="comment">	footer := []byte("</code></pre>\n\n_____")</span>
<span id="L145" class="line" lang="go"><a href="#L145">145</a>	</span><span class="comment"></span>
<span id="L146" class="line" lang="go"><a href="#L146">146</a>	</span><span class="comment">	buf = byteJoin(buf, footer)</span>
<span id="L147" class="line" lang="go"><a href="#L147">147</a>	</span><span class="comment"></span>
<span id="L148" class="line" lang="go"><a href="#L148">148</a>	</span><span class="comment">	s.Content = gen</span>
<span id="L149" class="line" lang="go"><a href="#L149">149</a>	</span><span class="comment">	return buf</span>
<span id="L150" class="line" lang="go"><a href="#L150">150</a>	</span><span class="comment"></span>
<span id="L151" class="line" lang="go"><a href="#L151">151</a>	</span><span class="comment">}</span>
<span id="L152" class="line" lang="go"><a href="#L152">152</a>	</span><span class="comment"></span>
<span id="L153" class="line" lang="go"><a href="#L153">153</a>	</span><span class="comment">func byteJoin(input ...[]byte) []byte {</span>
<span id="L154" class="line" lang="go"><a href="#L154">154</a>	</span><span class="comment">	var empty []byte</span>
<span id="L155" class="line" lang="go"><a href="#L155">155</a>	</span><span class="comment">	array := bytes.Join(input, empty)</span>
<span id="L156" class="line" lang="go"><a href="#L156">156</a>	</span><span class="comment">	return array</span>
<span id="L157" class="line" lang="go"><a href="#L157">157</a>	</span><span class="comment">}</span>
<span id="L158" class="line" lang="go"><a href="#L158">158</a>	</span><span class="comment"></span>
<span id="L159" class="line" lang="go"><a href="#L159">159</a>	</span><span class="comment">func mdURL(s, u string) string {</span>
<span id="L160" class="line" lang="go"><a href="#L160">160</a>	</span><span class="comment">	return "[" + s + "](" + u + ")"</span>
<span id="L161" class="line" lang="go"><a href="#L161">161</a>	</span><span class="comment">}</span>
<span id="L162" class="line" lang="go"><a href="#L162">162</a>	</span><span class="comment"></span>
<span id="L163" class="line" lang="go"><a href="#L163">163</a>	</span><span class="comment">func mdRelURL(s string, i int) string {</span>
<span id="L164" class="line" lang="go"><a href="#L164">164</a>	</span><span class="comment"></span>
<span id="L165" class="line" lang="go"><a href="#L165">165</a>	</span><span class="comment">	if i <= 0 {</span>
<span id="L166" class="line" lang="go"><a href="#L166">166</a>	</span><span class="comment">		return mdURL(s, "./"+s+".md")</span>
<span id="L167" class="line" lang="go"><a href="#L167">167</a>	</span><span class="comment"></span>
<span id="L168" class="line" lang="go"><a href="#L168">168</a>	</span><span class="comment">	}</span>
<span id="L169" class="line" lang="go"><a href="#L169">169</a>	</span><span class="comment">	var relPath string</span>
<span id="L170" class="line" lang="go"><a href="#L170">170</a>	</span><span class="comment"></span>
<span id="L171" class="line" lang="go"><a href="#L171">171</a>	</span><span class="comment">	for idx := 1; idx <= i; idx++ {</span>
<span id="L172" class="line" lang="go"><a href="#L172">172</a>	</span><span class="comment">		relPath += "../"</span>
<span id="L173" class="line" lang="go"><a href="#L173">173</a>	</span><span class="comment">	}</span>
<span id="L174" class="line" lang="go"><a href="#L174">174</a>	</span><span class="comment">	return mdURL(s, relPath+s+"/index.md")</span>
<span id="L175" class="line" lang="go"><a href="#L175">175</a>	</span><span class="comment"></span>
<span id="L176" class="line" lang="go"><a href="#L176">176</a>	</span><span class="comment">}</span>
</code></pre>

_____