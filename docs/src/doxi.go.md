## /[src](../src/index.md)/[doxi.go](./doxi.go)

<pre class="code highlight"><code>
<span id="L1" class="line" lang="go"><a href="#L1">1</a>	</span><span class="comment">// Doxi is a Go source code documentation generator, which generates static</span>
<span id="L2" class="line" lang="go"><a href="#L2">2</a>	</span><span class="comment">// Markdown / HTML files for you to host. The point is to generate a set of</span>
<span id="L3" class="line" lang="go"><a href="#L3">3</a>	</span><span class="comment">// documents similar to the Golang source code documentation, as seen with</span>
<span id="L4" class="line" lang="go"><a href="#L4">4</a>	</span><span class="comment">// `godoc` and `go doc`, however having a static Markdown / HTML render,</span>
<span id="L5" class="line" lang="go"><a href="#L5">5</a>	</span><span class="comment">// with relative hyperlinks instead of absolute ones.</span>
<span id="L6" class="line" lang="go"><a href="#L6">6</a>	</span><span class="comment">//</span>
<span id="L7" class="line" lang="go"><a href="#L7">7</a>	</span><span class="comment">// You can easily document your project by running the `doxi` binary, or by</span>
<span id="L8" class="line" lang="go"><a href="#L8">8</a>	</span><span class="comment">// embedding simple logic as below into your existing project, as a closer,</span>
<span id="L9" class="line" lang="go"><a href="#L9">9</a>	</span><span class="comment">// or "onDone" type of function:</span>
<span id="L10" class="line" lang="go"><a href="#L10">10</a>	</span><span class="comment">//</span>
<span id="L11" class="line" lang="go"><a href="#L11">11</a>	</span><span class="comment">//     package main</span>
<span id="L12" class="line" lang="go"><a href="#L12">12</a>	</span><span class="comment">//</span>
<span id="L13" class="line" lang="go"><a href="#L13">13</a>	</span><span class="comment">//     import (</span>
<span id="L14" class="line" lang="go"><a href="#L14">14</a>	</span><span class="comment">//	       "github.com/ZalgoNoise/doxi/dox"</span>
<span id="L15" class="line" lang="go"><a href="#L15">15</a>	</span><span class="comment">//     )</span>
<span id="L16" class="line" lang="go"><a href="#L16">16</a>	</span><span class="comment">//</span>
<span id="L17" class="line" lang="go"><a href="#L17">17</a>	</span><span class="comment">//     func main() {</span>
<span id="L18" class="line" lang="go"><a href="#L18">18</a>	</span><span class="comment">//         // your code goes here</span>
<span id="L19" class="line" lang="go"><a href="#L19">19</a>	</span><span class="comment">//         dox := dox.New()</span>
<span id="L20" class="line" lang="go"><a href="#L20">20</a>	</span><span class="comment">//         dox.Run()</span>
<span id="L21" class="line" lang="go"><a href="#L21">21</a>	</span><span class="comment">//     }</span>
<span id="L22" class="line" lang="go"><a href="#L22">22</a>	</span><span class="comment">//</span>
<span id="L23" class="line" lang="go"><a href="#L23">23</a>	</span><span class="comment">//</span>
<span id="L24" class="line" lang="go"><a href="#L24">24</a>	</span><span class="comment">package main</span>
<span id="L25" class="line" lang="go"><a href="#L25">25</a>	</span><span class="comment"></span>
<span id="L26" class="line" lang="go"><a href="#L26">26</a>	</span><span class="comment">import (</span>
<span id="L27" class="line" lang="go"><a href="#L27">27</a>	</span><span class="comment">	"github.com/ZalgoNoise/doxi/dox"</span>
<span id="L28" class="line" lang="go"><a href="#L28">28</a>	</span><span class="comment">)</span>
<span id="L29" class="line" lang="go"><a href="#L29">29</a>	</span><span class="comment"></span>
<span id="L30" class="line" lang="go"><a href="#L30">30</a>	</span><span class="comment">func main() {</span>
<span id="L31" class="line" lang="go"><a href="#L31">31</a>	</span><span class="comment">	dox := dox.New()</span>
<span id="L32" class="line" lang="go"><a href="#L32">32</a>	</span><span class="comment">	dox.Run()</span>
<span id="L33" class="line" lang="go"><a href="#L33">33</a>	</span><span class="comment"></span>
<span id="L34" class="line" lang="go"><a href="#L34">34</a>	</span><span class="comment">}</span>
</code></pre>

_____