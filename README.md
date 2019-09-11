# avatar-practice
**From the [Go Bootcamp](https://github.com/GoBootcamp/avatarme):** Given a personal information such as an email address, IP address, or a public key, the program you will write needs to generate a unique avatar. Imagine that you are building a new application and you want all of your users to have a default and unique avatar. The package you will write will allow the generation of such avatars. GitHub recently used such an approach and generates an identicon for all new users who don't have a gravatar account attached.

## Resources
http://golang.org/pkg/crypto/ <br>
http://golang.org/pkg/image/ <br>
http://en.wikipedia.org/wiki/Identicon <br>
http://haacked.com/archive/2007/01/22/Identicons_as_Visual_Fingerprints.aspx/

## Takeaways & challenges

After working on this for a while, I've learned that the main challenge of this exercise is to figure out the best logic for associating the values in the hash with characteristics of the desired image. I found this challenging and am still working on ways to improve it so that each image is truly unique. Right now, the uniqueness depends mainly on the colors in the image, which originate from some of the hash values. For instance, the following inputs have these results:

<img src="./samples/hello.png" alt="hello" height="200px" width="200px">

input: `hello` <br>

<img src="./samples/hellomy.png" alt="hellomy" height="200px" width="200px">

input: `hellomy` <br>

<img src="./samples/hellomyname.png" alt="hellomyname" height="200px" width="200px">

input: `hellomyname` <br>

<img src="./samples/hellomynameis.png" alt="hellomynameis" height="200px" width="200px">

input: `hellomynameis` <br>

<img src="./samples/hellomynameiszoe.png" alt="hellomynameiszoe" height="200px" width="200px">

input: `hellomynameiszoe` <br>

While a few of these are somewhat unique, the current code does not use all of the values from the hash, and this could be improved to have more unique results. I think more research needs to be done to see how this solution is implemented in real world situations. Until then...

<img src="./samples/thanksforvisitingmyrepo.png" alt="thanksforvisitingmyrepo" height="200px" width="200px">

` thanksforvisitingmyrepo ` :smile: