今天用IntelliJ IDEA导入Gradle项目时，一直提示没有配置JAVA_HOME,但是我确实是配置过这些环境变量，原来这里有些小坑。

**环境配置是：**

- Mac OS X 10.10.4 x86_64
- Java version "1.8.0_25"
- Gradle 2.5
- IntelliJ IDEA 14

首先，确认配置了JAVA_HOME和GRADLE_HOME：

打开.bash_profile文件，看是否存在这样的配置：

>JAVA_HOME=\`/usr/libexec/java_home`
>export JAVA_HOME
>
>GRADLE_HOME=/usr/local/bin/gradle-2.5
>export GRADLE_HOME
>export PATH=$PATH:$GRADLE_HOME/bin

打开shell，运行java -version和gradle -version，如果没有错误提示，则表示配置正确。

<!--more-->

接下来，打开/Library/Java/JavaVirtualMachines/jdk1.8.0_25.jdk/Contents/Info.plist（jdk1.8.0_25.jdk请根据实际版本替换），找到
<pre class="prettyprint lang-XML">
    &lt;key&gt;JVMCapabilities&lt;/key&gt;
      &lt;array&gt;
          &lt;string&gt;CommandLine&lt;/string&gt;
      &lt;/array&gt;
</pre>      
为其再添加四个string
<pre class="prettyprint">
    &lt;string&gt;JNI&lt;/string&gt;    
    &lt;string&gt;BundledApp&lt;/string&gt;  
    &lt;string&gt;WebStart&lt;/string&gt;  
    &lt;string&gt;Applets&lt;/string&gt;
</pre>    
然后，打开/Applications/IntelliJ IDEA 14.app/Contents/Info.plist文件（IntelliJ IDEA 14.app根请据实际版本替换），找到
<pre class="prettyprint">
    &lt;key&gt;JVMVersion&lt;/key&gt;
	&lt;string&gt;1.6*，1.7*&lt;/string&gt; 
</pre>	
将1.6*，1.7+改为1.8并保存。

接下来打开IntelliJ IDEA,点击右下角的Configure->Project Defaults->Project Structure->New->JDK，选择所在目录。

现在，IntelliJ IDEA可以正常使用了。