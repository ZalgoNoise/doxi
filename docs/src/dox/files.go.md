## /[src](../../src/index.md)/[dox](../dox/index.md)/[files.go](./files.go.md)

<pre class="code highlight"><code>
<span id="L1" class="line" lang="go"><a href="#L1">1</a>	</span><span class="comment">package dox</span>
<span id="L2" class="line" lang="go"><a href="#L2">2</a>	</span><span class="comment"></span>
<span id="L3" class="line" lang="go"><a href="#L3">3</a>	</span><span class="comment">import (</span>
<span id="L4" class="line" lang="go"><a href="#L4">4</a>	</span><span class="comment">	"os"</span>
<span id="L5" class="line" lang="go"><a href="#L5">5</a>	</span><span class="comment">	"strings"</span>
<span id="L6" class="line" lang="go"><a href="#L6">6</a>	</span><span class="comment">)</span>
<span id="L7" class="line" lang="go"><a href="#L7">7</a>	</span><span class="comment"></span>
<span id="L8" class="line" lang="go"><a href="#L8">8</a>	</span><span class="comment">// File struct will describe a file object, within the context of a</span>
<span id="L9" class="line" lang="go"><a href="#L9">9</a>	</span><span class="comment">// project. It is supposed to hold all required metadata to build</span>
<span id="L10" class="line" lang="go"><a href="#L10">10</a>	</span><span class="comment">// documentation with Doxi.</span>
<span id="L11" class="line" lang="go"><a href="#L11">11</a>	</span><span class="comment">type File struct {</span>
<span id="L12" class="line" lang="go"><a href="#L12">12</a>	</span><span class="comment">	Proj    Proj    `json:"project"`</span>
<span id="L13" class="line" lang="go"><a href="#L13">13</a>	</span><span class="comment">	Path    Path    `json:"abs_path"`</span>
<span id="L14" class="line" lang="go"><a href="#L14">14</a>	</span><span class="comment">	RelPath RelPath `json:"rel_path"`</span>
<span id="L15" class="line" lang="go"><a href="#L15">15</a>	</span><span class="comment">	Parent  Parent  `json:"parent_dir"`</span>
<span id="L16" class="line" lang="go"><a href="#L16">16</a>	</span><span class="comment">	Name    Name    `json:"file_name"`</span>
<span id="L17" class="line" lang="go"><a href="#L17">17</a>	</span><span class="comment">	Ext     Ext     `json:"extension"`</span>
<span id="L18" class="line" lang="go"><a href="#L18">18</a>	</span><span class="comment">	Type    Type    `json:"type"`</span>
<span id="L19" class="line" lang="go"><a href="#L19">19</a>	</span><span class="comment">	Source  *File   `json:"mem_ref"`</span>
<span id="L20" class="line" lang="go"><a href="#L20">20</a>	</span><span class="comment">}</span>
<span id="L21" class="line" lang="go"><a href="#L21">21</a>	</span><span class="comment"></span>
<span id="L22" class="line" lang="go"><a href="#L22">22</a>	</span><span class="comment">// Proj type will represent the project name (parent folder to current dir)</span>
<span id="L23" class="line" lang="go"><a href="#L23">23</a>	</span><span class="comment">type Proj string</span>
<span id="L24" class="line" lang="go"><a href="#L24">24</a>	</span><span class="comment"></span>
<span id="L25" class="line" lang="go"><a href="#L25">25</a>	</span><span class="comment">// Path is the absolute path to the project's directory</span>
<span id="L26" class="line" lang="go"><a href="#L26">26</a>	</span><span class="comment">type Path string</span>
<span id="L27" class="line" lang="go"><a href="#L27">27</a>	</span><span class="comment"></span>
<span id="L28" class="line" lang="go"><a href="#L28">28</a>	</span><span class="comment">// RelPath is the relative path within the project's directory</span>
<span id="L29" class="line" lang="go"><a href="#L29">29</a>	</span><span class="comment">type RelPath string</span>
<span id="L30" class="line" lang="go"><a href="#L30">30</a>	</span><span class="comment"></span>
<span id="L31" class="line" lang="go"><a href="#L31">31</a>	</span><span class="comment">// Parent is the parent folder for this file</span>
<span id="L32" class="line" lang="go"><a href="#L32">32</a>	</span><span class="comment">type Parent string</span>
<span id="L33" class="line" lang="go"><a href="#L33">33</a>	</span><span class="comment"></span>
<span id="L34" class="line" lang="go"><a href="#L34">34</a>	</span><span class="comment">// Name is the name of the file</span>
<span id="L35" class="line" lang="go"><a href="#L35">35</a>	</span><span class="comment">type Name string</span>
<span id="L36" class="line" lang="go"><a href="#L36">36</a>	</span><span class="comment"></span>
<span id="L37" class="line" lang="go"><a href="#L37">37</a>	</span><span class="comment">// Ext is the file's extension</span>
<span id="L38" class="line" lang="go"><a href="#L38">38</a>	</span><span class="comment">type Ext string</span>
<span id="L39" class="line" lang="go"><a href="#L39">39</a>	</span><span class="comment"></span>
<span id="L40" class="line" lang="go"><a href="#L40">40</a>	</span><span class="comment">// Type will describe if it's a file or folder (folders are files in Unix)</span>
<span id="L41" class="line" lang="go"><a href="#L41">41</a>	</span><span class="comment">type Type string</span>
<span id="L42" class="line" lang="go"><a href="#L42">42</a>	</span><span class="comment"></span>
<span id="L43" class="line" lang="go"><a href="#L43">43</a>	</span><span class="comment">// NewFile function will create a new instance of File</span>
<span id="L44" class="line" lang="go"><a href="#L44">44</a>	</span><span class="comment">func NewFile() *File {</span>
<span id="L45" class="line" lang="go"><a href="#L45">45</a>	</span><span class="comment">	f := &File{}</span>
<span id="L46" class="line" lang="go"><a href="#L46">46</a>	</span><span class="comment">	return f</span>
<span id="L47" class="line" lang="go"><a href="#L47">47</a>	</span><span class="comment">}</span>
<span id="L48" class="line" lang="go"><a href="#L48">48</a>	</span><span class="comment"></span>
<span id="L49" class="line" lang="go"><a href="#L49">49</a>	</span><span class="comment">// Init method will initialize a new file based on its (base) path and</span>
<span id="L50" class="line" lang="go"><a href="#L50">50</a>	</span><span class="comment">// the actual file path</span>
<span id="L51" class="line" lang="go"><a href="#L51">51</a>	</span><span class="comment">func (f *File) Init(p, s string) *File {</span>
<span id="L52" class="line" lang="go"><a href="#L52">52</a>	</span><span class="comment">	f.Path.Set(f, s)</span>
<span id="L53" class="line" lang="go"><a href="#L53">53</a>	</span><span class="comment"></span>
<span id="L54" class="line" lang="go"><a href="#L54">54</a>	</span><span class="comment">	proj := f.Proj.Gen(f, p)</span>
<span id="L55" class="line" lang="go"><a href="#L55">55</a>	</span><span class="comment">	f.Proj.Set(f, proj)</span>
<span id="L56" class="line" lang="go"><a href="#L56">56</a>	</span><span class="comment"></span>
<span id="L57" class="line" lang="go"><a href="#L57">57</a>	</span><span class="comment">	file := strings.Split(s, p)</span>
<span id="L58" class="line" lang="go"><a href="#L58">58</a>	</span><span class="comment">	if len(file) > 1 {</span>
<span id="L59" class="line" lang="go"><a href="#L59">59</a>	</span><span class="comment">		f.RelPath.Set(f, file[1])</span>
<span id="L60" class="line" lang="go"><a href="#L60">60</a>	</span><span class="comment">	}</span>
<span id="L61" class="line" lang="go"><a href="#L61">61</a>	</span><span class="comment"></span>
<span id="L62" class="line" lang="go"><a href="#L62">62</a>	</span><span class="comment">	parent := strings.Split(file[1], "/")</span>
<span id="L63" class="line" lang="go"><a href="#L63">63</a>	</span><span class="comment">	f.Parent.Set(f, parent[(len(parent)-2)])</span>
<span id="L64" class="line" lang="go"><a href="#L64">64</a>	</span><span class="comment"></span>
<span id="L65" class="line" lang="go"><a href="#L65">65</a>	</span><span class="comment">	f.Name.Set(f, parent[(len(parent)-1)])</span>
<span id="L66" class="line" lang="go"><a href="#L66">66</a>	</span><span class="comment"></span>
<span id="L67" class="line" lang="go"><a href="#L67">67</a>	</span><span class="comment">	ext := strings.Split(file[1], ".")</span>
<span id="L68" class="line" lang="go"><a href="#L68">68</a>	</span><span class="comment">	if len(ext) > 1 {</span>
<span id="L69" class="line" lang="go"><a href="#L69">69</a>	</span><span class="comment">		f.Ext.Set(f, "."+ext[(len(ext)-1)])</span>
<span id="L70" class="line" lang="go"><a href="#L70">70</a>	</span><span class="comment">		f.Type.Set(f, "file")</span>
<span id="L71" class="line" lang="go"><a href="#L71">71</a>	</span><span class="comment">	} else if _, err := os.Stat(s); !os.IsNotExist(err) {</span>
<span id="L72" class="line" lang="go"><a href="#L72">72</a>	</span><span class="comment">		f.Type.Set(f, "dir")</span>
<span id="L73" class="line" lang="go"><a href="#L73">73</a>	</span><span class="comment">	}</span>
<span id="L74" class="line" lang="go"><a href="#L74">74</a>	</span><span class="comment"></span>
<span id="L75" class="line" lang="go"><a href="#L75">75</a>	</span><span class="comment">	return f</span>
<span id="L76" class="line" lang="go"><a href="#L76">76</a>	</span><span class="comment">}</span>
<span id="L77" class="line" lang="go"><a href="#L77">77</a>	</span><span class="comment"></span>
<span id="L78" class="line" lang="go"><a href="#L78">78</a>	</span><span class="comment">// Unite method will create mutual symlinks between File objects</span>
<span id="L79" class="line" lang="go"><a href="#L79">79</a>	</span><span class="comment">// (memory-wise)</span>
<span id="L80" class="line" lang="go"><a href="#L80">80</a>	</span><span class="comment">func (f *File) Unite(i File) File {</span>
<span id="L81" class="line" lang="go"><a href="#L81">81</a>	</span><span class="comment">	f.Source = &i</span>
<span id="L82" class="line" lang="go"><a href="#L82">82</a>	</span><span class="comment">	i.Source = f</span>
<span id="L83" class="line" lang="go"><a href="#L83">83</a>	</span><span class="comment">	return i</span>
<span id="L84" class="line" lang="go"><a href="#L84">84</a>	</span><span class="comment"></span>
<span id="L85" class="line" lang="go"><a href="#L85">85</a>	</span><span class="comment">}</span>
<span id="L86" class="line" lang="go"><a href="#L86">86</a>	</span><span class="comment"></span>
<span id="L87" class="line" lang="go"><a href="#L87">87</a>	</span><span class="comment">// Fetch method will return the source / origin / link of a requested file</span>
<span id="L88" class="line" lang="go"><a href="#L88">88</a>	</span><span class="comment">func (f *File) Fetch() File {</span>
<span id="L89" class="line" lang="go"><a href="#L89">89</a>	</span><span class="comment">	return *f.Source</span>
<span id="L90" class="line" lang="go"><a href="#L90">90</a>	</span><span class="comment">}</span>
<span id="L91" class="line" lang="go"><a href="#L91">91</a>	</span><span class="comment"></span>
<span id="L92" class="line" lang="go"><a href="#L92">92</a>	</span><span class="comment">// Link method will create a memory symlink with both File objects</span>
<span id="L93" class="line" lang="go"><a href="#L93">93</a>	</span><span class="comment">func (f *File) Link(i File) {</span>
<span id="L94" class="line" lang="go"><a href="#L94">94</a>	</span><span class="comment">	f.Source = &i</span>
<span id="L95" class="line" lang="go"><a href="#L95">95</a>	</span><span class="comment">}</span>
<span id="L96" class="line" lang="go"><a href="#L96">96</a>	</span><span class="comment"></span>
<span id="L97" class="line" lang="go"><a href="#L97">97</a>	</span><span class="comment">// Gen method - Proj - Generate the reference project name based on the</span>
<span id="L98" class="line" lang="go"><a href="#L98">98</a>	</span><span class="comment">// current file's relative path. If the object derives from a Docs object,</span>
<span id="L99" class="line" lang="go"><a href="#L99">99</a>	</span><span class="comment">// jump one folder up; otherwise it's the last "entry" in the path</span>
<span id="L100" class="line" lang="go"><a href="#L100">100</a>	</span><span class="comment">func (f *Proj) Gen(i *File, s string) string {</span>
<span id="L101" class="line" lang="go"><a href="#L101">101</a>	</span><span class="comment"></span>
<span id="L102" class="line" lang="go"><a href="#L102">102</a>	</span><span class="comment">	path := strings.Split(s, "/")</span>
<span id="L103" class="line" lang="go"><a href="#L103">103</a>	</span><span class="comment"></span>
<span id="L104" class="line" lang="go"><a href="#L104">104</a>	</span><span class="comment">	if path[(len(path)-1)] == "docs" {</span>
<span id="L105" class="line" lang="go"><a href="#L105">105</a>	</span><span class="comment">		return path[(len(path) - 2)]</span>
<span id="L106" class="line" lang="go"><a href="#L106">106</a>	</span><span class="comment">	}</span>
<span id="L107" class="line" lang="go"><a href="#L107">107</a>	</span><span class="comment">	return path[(len(path) - 1)]</span>
<span id="L108" class="line" lang="go"><a href="#L108">108</a>	</span><span class="comment">}</span>
<span id="L109" class="line" lang="go"><a href="#L109">109</a>	</span><span class="comment"></span>
<span id="L110" class="line" lang="go"><a href="#L110">110</a>	</span><span class="comment">// Set method - Proj - defines this type with an input File and string</span>
<span id="L111" class="line" lang="go"><a href="#L111">111</a>	</span><span class="comment">func (f *Proj) Set(i *File, s string) {</span>
<span id="L112" class="line" lang="go"><a href="#L112">112</a>	</span><span class="comment">	new := Proj(s)</span>
<span id="L113" class="line" lang="go"><a href="#L113">113</a>	</span><span class="comment">	i.Proj = new</span>
<span id="L114" class="line" lang="go"><a href="#L114">114</a>	</span><span class="comment">}</span>
<span id="L115" class="line" lang="go"><a href="#L115">115</a>	</span><span class="comment"></span>
<span id="L116" class="line" lang="go"><a href="#L116">116</a>	</span><span class="comment">// Set method - Path - defines this type with an input File and string</span>
<span id="L117" class="line" lang="go"><a href="#L117">117</a>	</span><span class="comment">func (f *Path) Set(i *File, s string) {</span>
<span id="L118" class="line" lang="go"><a href="#L118">118</a>	</span><span class="comment">	new := Path(s)</span>
<span id="L119" class="line" lang="go"><a href="#L119">119</a>	</span><span class="comment">	i.Path = new</span>
<span id="L120" class="line" lang="go"><a href="#L120">120</a>	</span><span class="comment">}</span>
<span id="L121" class="line" lang="go"><a href="#L121">121</a>	</span><span class="comment"></span>
<span id="L122" class="line" lang="go"><a href="#L122">122</a>	</span><span class="comment">// Set method - RelPath - defines this type with an input File and string</span>
<span id="L123" class="line" lang="go"><a href="#L123">123</a>	</span><span class="comment">func (f *RelPath) Set(i *File, s string) {</span>
<span id="L124" class="line" lang="go"><a href="#L124">124</a>	</span><span class="comment">	new := RelPath(s)</span>
<span id="L125" class="line" lang="go"><a href="#L125">125</a>	</span><span class="comment">	i.RelPath = new</span>
<span id="L126" class="line" lang="go"><a href="#L126">126</a>	</span><span class="comment">}</span>
<span id="L127" class="line" lang="go"><a href="#L127">127</a>	</span><span class="comment"></span>
<span id="L128" class="line" lang="go"><a href="#L128">128</a>	</span><span class="comment">// Set method - Parent - defines this type with an input File and string</span>
<span id="L129" class="line" lang="go"><a href="#L129">129</a>	</span><span class="comment">func (f *Parent) Set(i *File, s string) {</span>
<span id="L130" class="line" lang="go"><a href="#L130">130</a>	</span><span class="comment">	new := Parent(s)</span>
<span id="L131" class="line" lang="go"><a href="#L131">131</a>	</span><span class="comment">	i.Parent = new</span>
<span id="L132" class="line" lang="go"><a href="#L132">132</a>	</span><span class="comment">}</span>
<span id="L133" class="line" lang="go"><a href="#L133">133</a>	</span><span class="comment"></span>
<span id="L134" class="line" lang="go"><a href="#L134">134</a>	</span><span class="comment">// Set method - Name - defines this type with an input File and string</span>
<span id="L135" class="line" lang="go"><a href="#L135">135</a>	</span><span class="comment">func (f *Name) Set(i *File, s string) {</span>
<span id="L136" class="line" lang="go"><a href="#L136">136</a>	</span><span class="comment">	new := Name(s)</span>
<span id="L137" class="line" lang="go"><a href="#L137">137</a>	</span><span class="comment">	i.Name = new</span>
<span id="L138" class="line" lang="go"><a href="#L138">138</a>	</span><span class="comment">}</span>
<span id="L139" class="line" lang="go"><a href="#L139">139</a>	</span><span class="comment"></span>
<span id="L140" class="line" lang="go"><a href="#L140">140</a>	</span><span class="comment">// Set method - Ext - defines this type with an input File and string</span>
<span id="L141" class="line" lang="go"><a href="#L141">141</a>	</span><span class="comment">func (f *Ext) Set(i *File, s string) {</span>
<span id="L142" class="line" lang="go"><a href="#L142">142</a>	</span><span class="comment">	new := Ext(s)</span>
<span id="L143" class="line" lang="go"><a href="#L143">143</a>	</span><span class="comment">	i.Ext = new</span>
<span id="L144" class="line" lang="go"><a href="#L144">144</a>	</span><span class="comment">}</span>
<span id="L145" class="line" lang="go"><a href="#L145">145</a>	</span><span class="comment"></span>
<span id="L146" class="line" lang="go"><a href="#L146">146</a>	</span><span class="comment">// Set method - Type - defines this type with an input File and string</span>
<span id="L147" class="line" lang="go"><a href="#L147">147</a>	</span><span class="comment">func (f *Type) Set(i *File, s string) {</span>
<span id="L148" class="line" lang="go"><a href="#L148">148</a>	</span><span class="comment">	new := Type(s)</span>
<span id="L149" class="line" lang="go"><a href="#L149">149</a>	</span><span class="comment">	i.Type = new</span>
<span id="L150" class="line" lang="go"><a href="#L150">150</a>	</span><span class="comment">}</span>
<span id="L151" class="line" lang="go"><a href="#L151">151</a>	</span><span class="comment"></span>
<span id="L152" class="line" lang="go"><a href="#L152">152</a>	</span><span class="comment">// Get method - Proj - retrieves the string value for this type</span>
<span id="L153" class="line" lang="go"><a href="#L153">153</a>	</span><span class="comment">func (f *Proj) Get() string {</span>
<span id="L154" class="line" lang="go"><a href="#L154">154</a>	</span><span class="comment">	return string(*f)</span>
<span id="L155" class="line" lang="go"><a href="#L155">155</a>	</span><span class="comment">}</span>
<span id="L156" class="line" lang="go"><a href="#L156">156</a>	</span><span class="comment"></span>
<span id="L157" class="line" lang="go"><a href="#L157">157</a>	</span><span class="comment">// Get method - Path - retrieves the string value for this type</span>
<span id="L158" class="line" lang="go"><a href="#L158">158</a>	</span><span class="comment">func (f *Path) Get() string {</span>
<span id="L159" class="line" lang="go"><a href="#L159">159</a>	</span><span class="comment">	return string(*f)</span>
<span id="L160" class="line" lang="go"><a href="#L160">160</a>	</span><span class="comment">}</span>
<span id="L161" class="line" lang="go"><a href="#L161">161</a>	</span><span class="comment"></span>
<span id="L162" class="line" lang="go"><a href="#L162">162</a>	</span><span class="comment">// Get method - RelPath - retrieves the string value for this type</span>
<span id="L163" class="line" lang="go"><a href="#L163">163</a>	</span><span class="comment">func (f *RelPath) Get() string {</span>
<span id="L164" class="line" lang="go"><a href="#L164">164</a>	</span><span class="comment">	return string(*f)</span>
<span id="L165" class="line" lang="go"><a href="#L165">165</a>	</span><span class="comment">}</span>
<span id="L166" class="line" lang="go"><a href="#L166">166</a>	</span><span class="comment"></span>
<span id="L167" class="line" lang="go"><a href="#L167">167</a>	</span><span class="comment">// Get method - Parent - retrieves the string value for this type</span>
<span id="L168" class="line" lang="go"><a href="#L168">168</a>	</span><span class="comment">func (f *Parent) Get() string {</span>
<span id="L169" class="line" lang="go"><a href="#L169">169</a>	</span><span class="comment">	return string(*f)</span>
<span id="L170" class="line" lang="go"><a href="#L170">170</a>	</span><span class="comment">}</span>
<span id="L171" class="line" lang="go"><a href="#L171">171</a>	</span><span class="comment"></span>
<span id="L172" class="line" lang="go"><a href="#L172">172</a>	</span><span class="comment">// Get method - Name - retrieves the string value for this type</span>
<span id="L173" class="line" lang="go"><a href="#L173">173</a>	</span><span class="comment">func (f *Name) Get() string {</span>
<span id="L174" class="line" lang="go"><a href="#L174">174</a>	</span><span class="comment">	return string(*f)</span>
<span id="L175" class="line" lang="go"><a href="#L175">175</a>	</span><span class="comment">}</span>
<span id="L176" class="line" lang="go"><a href="#L176">176</a>	</span><span class="comment"></span>
<span id="L177" class="line" lang="go"><a href="#L177">177</a>	</span><span class="comment">// Get method - Ext - retrieves the string value for this type</span>
<span id="L178" class="line" lang="go"><a href="#L178">178</a>	</span><span class="comment">func (f *Ext) Get() string {</span>
<span id="L179" class="line" lang="go"><a href="#L179">179</a>	</span><span class="comment">	return string(*f)</span>
<span id="L180" class="line" lang="go"><a href="#L180">180</a>	</span><span class="comment">}</span>
<span id="L181" class="line" lang="go"><a href="#L181">181</a>	</span><span class="comment"></span>
<span id="L182" class="line" lang="go"><a href="#L182">182</a>	</span><span class="comment">// Get method - Type - retrieves the string value for this type</span>
<span id="L183" class="line" lang="go"><a href="#L183">183</a>	</span><span class="comment">func (f *Type) Get() string {</span>
<span id="L184" class="line" lang="go"><a href="#L184">184</a>	</span><span class="comment">	return string(*f)</span>
<span id="L185" class="line" lang="go"><a href="#L185">185</a>	</span><span class="comment">}</span>
</code></pre>

_____