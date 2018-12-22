### __DESCRIPTION__ 🤖🖥

Similar to BotPress and similar bot building platforms, the aim of this framework is to have a decoupled system where any additional features are implemented in the form of modules. Go-jack is heavily inspired by [BotPress](https://github.com/botpress/botpress)

#

#### Features 😎
1. Allow the creation of multi-channel bots

2. Seamless integration of NLP into the core, to more easily handle NLP in bots build using go-jack.

3. Support the use of API calls/webhooks to monitor/use the bot outside of go-jack (for example, in another Go project entirely without adding go-jack as another dependency in the project)

4. Suitable logging system (as bot error logs have to be as clear as they can be)

5. Better content management for non-technical bot builders --separate code from content to make it easier for non-technical co-builders/content creators add content without being concerned about the code or messing the codebase up.

6. More features as deem fit by the developer (to be expanded as time goes on)

#

### __Why Go-Jack__ 🤔

Not seeing suitable bot building frameworks in Golang. While there are tools/frameworks like [Go-sarah](https://github.com/oklahomer/go-sarah), they can be quite technical and even not so easy/flexible to pick up for technical people/coders.

There are also special frameworks targeted at building for specific platforms. But none suitable for building for multiple channels at once, *__separating content from code, and easy to pick up even for non-coders who need to create contents for their bots__*
#

### __What has changed__ ❓
Well, after much thinking, I decided to write the messenger module first. The idea is to write each platform separately starting with Messenger (most conversant), then working out the core that binds them together. In the end, the already written "frameworks" for each channel will be turned into modules. This is why this branch is called *__messenger/poc__*

A major reason I considered this approach is that:
1. I am having issues figuring out the module system in Go (I've only been writing Go for about 6-7 months.)
2. I want a separation of concern. Messenger platform has always been my point of interest among all channels (I even work on Messenger bots at my work)


After much digging, I have a better insight on how to structure and organize the messenger channel framework first. While a whole lot of codes already written in [Goblin](https://github.com/oayomide/goblin) already work and stable and quite standard, huge props to the [Messenger](https://github.com/paked/messenger) project and the contributors. It not only reinforces the codes I had already written, it serves as a source of inspiriation and ideas. This library is hugely a replica of the project with distinct features.

#

### __Contributing 🖊__

Of course the project sounds interesting and all that, but I may not get around to it for a really long time (as I am busy and I also have to work on other things apart from this). Therefore, contributions are highly welcome and will be appreciated.
#

### TO-DO 📓

1. [ ] Map out the architecture/middleware system
2. [x] Design basic logging system
3. [ ] Develop CLI
4. [ ] Develop basisc server

_If any of the above spark any joy and you'd love to dedicate time to this specially (beyond just few PRs), I'd love to make you (a) maintainer. You can simply say Hi privately 😄_