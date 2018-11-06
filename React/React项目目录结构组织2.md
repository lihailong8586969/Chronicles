> 原文：[How to better organize your React applications?](https://medium.com/@alexmngn/how-to-better-organize-your-react-applications-2fd3ea1920f1)
> 
> 在过去的几年，我一直在开发超大型的 Web 应用。与数十名开发者从零开始，使其现在达到服务数百万用户的规模。如果一开始就没有一个好的文件目录结构，就很难维持代码的组织性。
> 
> [Nathanael Beisiegel](http://engineering.kapost.com/author/nathanaelbeisiegel/) 写了一篇有趣的文章，阐述了他在组织大型 React 项目的策略。但我对其想法并不十分赞同。因此，我决定花些时间找出将来组织 React 项目的最佳方式。
> 
> 注意 1：在本文提及的所有案例均使用了 Redux。如果你还未了解 Redux，可以阅读 [此文档](http://redux.js.org/)。
> 
> 注意 2：所有案例均基于 ReactJS，但你可以在 React-Native 应用上使用相同的结构。
> 
> 写于 2018.04：如果想改善代码结构，那么你应该对我最近写的这篇文章感兴趣——[《为何 React 开发者应该懂得模块化》](https://medium.com/@alexmngn/why-react-developers-should-modularize-their-applications-d26d381854c1)
> 
> ### 构建程序时会遇到什么挑战？
> 绝大多数开发者会（或将会）碰到以下场景：
> 
> 1. 与团队数个开发者编写一个应用，期间一切工作良好。
> 2. 客户要求添加新功能。
> 3. 客户要求删除一些功能，并添加一个新功能。程序开始变得复杂，但你淡然置之。程序即使不完美，但仍然可以运行。
> 4. 客户再次要求更改和删除一些功能，并添加一个超出预期的功能。此时，你拿起透明胶带开始修补程序。当然，你并不以此为豪，尽管依然能运行。
> 5. 在接下来的 6 个月内，程序又经过数次迭代。代码的可读性变得非常差，犹如意大利面条。
> 
> 每当客户决定向程序添加新功能时，新代码都会与复杂的遗留代码混合在一起，使得可维护性越来越低。而导致这一切问题的根源在于**程序从一开始就没有正确设计**。
> 
> 在刚开始学习 React 时， ，我找了几篇关于如何创建 Todo List 或简单游戏的优秀文章。这些文章让我很好地理解了 React 的基础知识。但与此同时，我很快发现这并没有告诉我如何使用 React 构建真实项目，即包含数十个页面和数百个组件的项目。
> 
> 在搜索过后，我发现 Github 上所有 React 模板项目均是相类似的结构——根据**文件类型**组织项目。大概如下：
> 
> ```
> /src
>   /actions
>     /notifications.js
>       
>  /components 
>     /Header
>     /Footer
>     /Notifications
>       /index.js
>   /containers
>     /Home
>     /Login
>     /Notifications
>       /index.js
>   /images
>     /logo.png
>   /reducers 
>     /login.js
>     /notifications.js
>   /styles 
>     /app.scss
>     /header.scss 
>     /home.scss
>     /footer.scss
>     /notifications.scss
>   /utils
>   index.js  
> ```
> 这样的文件结构对于构建个人网站或应用应该没多大的问题，但我相信这并不是最佳的目录结构。
> 
> 当以文件类型组织时，无疑会随着应用的迭代而变得愈发难以维护。当你意识到这一点时已为时已晚，你不得不花费大量的时间和金钱去更改所有东西，或者在接下来的时间忍受着这一问题。
> 
> 使用 React 的好处是你可以随心所欲地组织项目，而不会强制你按照特定的文件结构。毕竟 React 就是一个简单的 JavaScript 库。
> 
> ### 何为更佳的应用组织方式？
> 在金融机构工作的那几年，我们将 Ember 作为主要的 JavaScript 框架去构建所有新的 Web 应用。而 Ember 有趣的一点是：能够按照功能组织项目，而不是文件类型。但这足以改变这一切。
> 
> Ember 的 [Pods](https://ember-cli.com/user-guide/#using-pods) 结构很优秀，但仍有局限性，所以我想为其赋予更多的灵活性。在经过数次实验以试图寻找出最佳的目录结构后，我得出这一结论：将所有相关功能组合在一起，然后按需嵌套。以下就是我正在使用的结构：
> 
> ```
> /src
>   /components 
>     /Button 
>     /Notifications
>       /components
>         /ButtonDismiss  
>           /images
>           /locales
>           /specs 
>           /index.js
>           /styles.scss
>       /index.js
>       /styles.scss
>   /scenes
>     /Home 
>       /components 
>         /ButtonLike
>       /services
>         /processData
>       /index.js
>       /styles.scss
>     /Sign 
>       /components 
>         /FormField
>       /scenes
>         /Login
>         /Register 
>           /locales
>           /specs
>           /index.js
>           /styles.scss
>   /services
>     /api
>     /geolocation
>     /session
>       /actions.js
>       /index.js
>       /reducer.js
>     /users
>       /actions.js
>       /api.js
>       /reducer.js
>   index.js 
>   store.js
> ```
> 每个组件、场景（译者注：一般为页面，但也可将登陆和注册两个页面视为一个场景）或服务（功能）均含有供自己使用的所有东西，如自身的样式、图片、译文（译者注：国际化）、actions（译者注：应该指 Redux 相关）和单元/集成测试。即将一项功能视为应用中的一个独立代码段（有点类似 Node 模块）。
> 
> 为了运行良好，需遵循以下规则：
> 
> 1. 组件内能定义嵌套组件或服务，但不能使用或定义场景。
> 2. 场景内能定义嵌套组件、场景或服务。
> 3. 服务内能定义嵌套服务，但不能使用或定义组件和场景。
> 4. 嵌套的功能只能归其父级使用。
> 
> 注意：父级不局限于直系父级，还包含祖先级。不能使用“堂兄弟”的功能，只能将其移到父级内再使用。
> 
> 下面听我分点阐述。
> 
> #### 组件（Components）
> 大家都知道组件是什么。但在此文件结构中，是允许组件嵌套组件。
> 
> 在项目根目录的 components 文件夹内定义的组件是全局的，即可在应用中的任意位置使用。当在一个组件内定义一个新组件（即嵌套）时，则该新组件就仅能被其直系父级使用。
> 
> ##### 为何这么做？
> 当开发一个大型应用时，你经常需要创建组件。你明确知道它不会在其他地方复用，但确实需要它。如果你将其放到根目录的 components 文件夹下，久而久之，里面就会填满数百个组件。此时，你意识到需要将它们进行分类，但早已忘了其用途和在哪些地方使用。
> 
> 如果仅在里面放应用的主要组件，如 button、form field、thumbnail 甚至是一些更复杂的组件，如 listComment、formComposer 及其子组件，那么就能提高查找它们的效率。
> 
> 案例：
> 
> ```
> /src
>   /components
>     /Button
>       /index.js
>     /Notifications 
>       /components 
>         /ButtonDismiss 
>           /index.js
>       /actions.js
>       /index.js
>       /reducer.js
> ```
> * Button 可在应用任意地方使用
> * Notifications 也可在应用任意地方使用。该组件内还定义了 ButtonDismiss 组件。但除了在 Notifications 内，其余地方均不能使用它。
> * ButtonDismiss 内使用了 Button。这是允许的，因为 Button 在根目录 components 内定义。
> 
> #### 场景（Scenes）
> 场景就是应用中的一个页面。场景其实也是一个组件，但我更乐意将它们放到单独一个文件夹中。
> 
> 如果有使用 [React-Router](https://github.com/reactjs/react-router) 或 [React Native Router](https://github.com/aksonov/react-native-router-flux)，那么你可以将所有场景引入到 index.js 应用入口文件中，并将它们与路由进行匹配。
> 
> 既然是组件，那么场景也是可嵌套的，即场景内嵌套场景。当然，也可以在场景内定义组件或服务。但需要记住的是：场景内定义的东西，仅供该场景使用。
> 
> 案例：
> 
> ```
> /src
>   /scenes
>     /Home 
>       /components
>         /ButtonShare
>           /index.js
>       /index.js
>     /Sign
>       /components
>         /ButtonHelp
>           /index.js
>       /scenes
>         /Login
>           /components 
>             /Form
>               /index.js
>             /ButtonFacebookLogin
>               /index.js
>           /index.js
>        
>         /Register
>           /index.js
>       /index.js
> ```
> * Home 拥有 ButtonShare 组件，但该组件仅供 Home 场景使用。
> * Sign 拥有 ButtonHelp 组件，该组件可供 Login 和 Register 场景使用。当然，也供它们内部定义的组件使用。
> * Form 组件内部使用了 ButtonHelp 组件。这是允许的，因为 ButtonHelp 由父级定义。
> * Register 场景不能使用在 Login 内定义的组件，但可以使用 ButtonHelp 组件。
> 
> #### 服务（Services）
> 并不是所有东西都能作为组件。你需要创建一个独立的模块，以供组件和场景使用。
> 
> 你可以将一项服务看作是独立的模块。它定义了应用的核心业务逻辑，最终可供多个场景，甚至是多个应用（web 与 native 版本 的应用）共享使用。
> 
> ```
> /src
>   /services
>     /api
>       /services
>         /handleError
>           /index.js
>       /index.js
>     /geolocation 
>     /session 
>       /actions.js
>       /index.js
>       /reducer.js
> ```
> 我建议你创建一个服务去管理所有 API 请求。即将其作为服务器 API 与视图层（场景和组件）的桥梁/适配器。它可以处理应用发起的网络请求、获取或发布内容、在发送或存储在应用的 store（如 Redux） 前按需处理负载。而场景和组件只需 dispatch actions、读取 store 数据和更新视图即可。
> 
> ### 总结
> 在过去几个月，我已经根据此目录结构开发了一个 React-Native 个人项目，并因此节省了大量时间。所以，将所有相关实体放在一起会让事件变得更简单和更易处理。
> 
> 这种目录结构只是无数种组织项目的方式之一。但这就是我现在喜欢的方式，希望它能提升你的工作效率！
> 
> 如果你想查看使用该文件结构的实际工作项目，可查看我 Github 账号：
> 
> * https://github.com/alexmngn/react-rock-paper-scissors (ReactJS)
> * https://github.com/alexmngn/react-native-authentication (React-Native)
> 
> 若有任何疑问，请随时评论或直接与我联系，我将非常乐意提供帮助。
> 
> Have fun!
> 
> #### 我的其他文章
> * [How to better organize your React applications?](https://medium.com/@alexmngn/how-to-better-organize-your-react-applications-2fd3ea1920f1)
> * [How to use Redux on highly scalable javascript applications?](https://medium.com/@alexmngn/how-to-use-redux-on-highly-scalable-javascript-applications-4e4b8cb5ef38)
> * [What are the main differences between ReactJS and React-Native?](https://medium.com/@alexmngn/from-reactjs-to-react-native-what-are-the-main-differences-between-both-d6e8e88ebf24)
> 
> #### 关于我
> 大家好，我是 Alexis！我是一名拥有 14 年开发经验的全栈开发者。我自小就对技术充满激情。现专注于 JavaScript 开发。乐意花费无数时间去学习和使用新技术，并将其用于我们下一个项目中。我最近开始使用 Twitter，欢迎关注我：[@alexmngn](https://twitter.com/alexmngn)

