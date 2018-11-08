# React ���˼��

> �����򣺱����� React ���Ŀ����ߡ��� React API �ս���֮�Ƶ� Sebastian Markb?ge ׫д������������� React �ĳ��ԡ��Ķ����ģ�����վ�ڸ��ߵĸ߶�˼�� React �Ĺ�ȥ�����ں�δ����ԭ�ĵ�ַ��https://github.com/reactjs/react-basic

��д����������ʽ�ز��������� React ��[����ģ��](http://baike.baidu.com/view/2333986.htm)��Ŀ���ǽ���Ϊʲô���ǻ�������� React��ͬʱ��Ҳ���Ը�����Щ�۵㷴�Ƴ� React��

���ɷ��ϣ������еĲ����۾ݻ�ǰ���д����飬���Ҳ���ʾ������ƿ��ܴ��� bug ���������ֻ����ʽȷ����������׶Ρ�������и��õ����������뷨������ʱ�ύ pull request�����Ĳ�����ܿ��ϸ���е��漼���ɣ�����������������죬���㿴�� React �ɼ򵥵����ӵ���ƹ��̡�

React.js ����ʵʵ���г����˾�������Ľ������������ʽ�Ľⷨ���㷨�Ż�����ʷ�������룬debug �����Լ�����һЩ����������ľ��и߿����Ե����ݡ���Щ������ܲ����ȶ�����Ϊδ��������ı仯�͹���Ȩ�صı仯��ʱ���ٸı䡣���Ծ���Ĵ�����ѳ��׽��������

��ƫ����ѡ��һ��������ȫ hold ס�ļ�������ģ���������ܡ�

## �任��Transformation��

��� React �ĺ���ǰ������Ϊ UI ֻ�ǰ�����ͨ��ӳ���ϵ�任����һ����ʽ�����ݡ�ͬ��������ػ���ͬ�����������ǡ�þ��Ǵ�������

```js
function NameBox(name) {
  return { fontWeight: 'bold', labelContent: name };
}
```

```
'Sebastian Markb?ge' ->
{ fontWeight: 'bold', labelContent: 'Sebastian Markb?ge' };
```

## ����Abstraction��

�㲻���ܽ���һ����������ʵ�ָ��ӵ� UI����Ҫ���ǣ�����Ҫ�� UI ����ɶ�������ڲ�ϸ�ڣ��ֿɸ��õĺ�����ͨ����һ�������е�����һ��������ʵ�ָ��ӵ� UI������ǳ���

```js
function FancyUserBox(user) {
  return {
    borderStyle: '1px solid blue',
    childContent: [
      'Name: ',
      NameBox(user.firstName + ' ' + user.lastName)
    ]
  };
}
```

```
{ firstName: 'Sebastian', lastName: 'Markb?ge' } ->
{
  borderStyle: '1px solid blue',
  childContent: [
    'Name: ',
    { fontWeight: 'bold', labelContent: 'Sebastian Markb?ge' }
  ]
};
```

## ��ϣ�Composition��

Ϊ�������ﵽ���õ����ԣ�ֻ����Ҷ��Ȼ��ÿ�ζ�Ϊ���Ǵ���һ���µ������ǲ����ġ��㻹��Ҫ���԰�����������������ٴν�����ϡ������ġ���ϡ����ǽ��������߶����ͬ�ĳ���ϲ�Ϊһ����

```js
function FancyBox(children) {
  return {
    borderStyle: '1px solid blue',
    children: children
  };
}

function UserBox(user) {
  return FancyBox([
    'Name: ',
    NameBox(user.firstName + ' ' + user.lastName)
  ]);
}
```

## ״̬��State��

UI �������ǶԷ������˻�ҵ���߼�״̬�ĸ��ơ�ʵ���ϻ��кܶ�״̬����Ծ������ȾĿ�ꡣ�ٸ����ӣ���һ�� text field �д��֡�����һ��Ҫ���Ƶ�����ҳ���������ֻ��豸������λ�����״̬��һ�����͵��㼸�����Ḵ�Ƶ������ȾĿ��ġ�

����������ʹ�ò��ɱ������ģ�͡����ǰѿ��Ըı� state �ĺ�������������Ϊԭ������ڶ��㡣

```js
function FancyNameBox(user, likes, onClick) {
  return FancyBox([
    'Name: ', NameBox(user.firstName + ' ' + user.lastName),
    'Likes: ', LikeBox(likes),
    LikeButton(onClick)
  ]);
}

// ʵ��ϸ��

var likes = 0;
function addOneMoreLike() {
  likes++;
  rerender();
}

// ��ʼ��

FancyNameBox(
  { firstName: 'Sebastian', lastName: 'Markb?ge' },
  likes,
  addOneMoreLike
);
```

*ע�⣺��������״̬ʱ����������ã�addOneMoreLike �����У�����ʵ�ʵ��뷨�ǵ�һ����update������ʱ���Ƿ�����һ���汾��״̬����������Ƚϸ��ӡ���ʾ��������*

## Memoization

���ڴ�������ʹ����ͬ�Ĳ���һ�δε���δ��̫�˷���Դ�����ǿ��Դ���һ�������� memorized �汾������׷�����һ�������ͽ��������������Ǽ���ʹ��ͬ����ֵ���Ͳ���Ҫ����ִ�����ˡ�

```js
function memoize(fn) {
  var cachedArg;
  var cachedResult;
  return function(arg) {
    if (cachedArg === arg) {
      return cachedResult;
    }
    cachedArg = arg;
    cachedResult = fn(arg);
    return cachedResult;
  };
}

var MemoizedNameBox = memoize(NameBox);

function NameAndAgeBox(user, currentTime) {
  return FancyBox([
    'Name: ',
    MemoizedNameBox(user.firstName + ' ' + user.lastName),
    'Age in milliseconds: ',
    currentTime - user.dateOfBirth
  ]);
}
```

## �б�Lists��

�󲿷� UI ����չʾ�б������в�ͬ item ���б�ṹ������һ����Ȼ�Ĳ㼶��

Ϊ�˹����б��е�ÿһ�� item �� state �����ǿ��Դ���һ�� Map ���ɾ��� item �� state��

```js
function UserList(users, likesPerUser, updateUserLikes) {
  return users.map(user => FancyNameBox(
    user,
    likesPerUser.get(user.id),
    () => updateUserLikes(user.id, likesPerUser.get(user.id) + 1)
  ));
}

var likesPerUser = new Map();
function updateUserLikes(id, likeCount) {
  likesPerUser.set(id, likeCount);
  rerender();
}

UserList(data.users, likesPerUser, updateUserLikes);
```

*ע�⣺���������� FancyNameBox ���˶����ͬ�Ĳ���������������ǵ� memoization ��Ϊ����ÿ��ֻ�ܴ洢һ��ֵ������������������档*

## �����ԣ�Continuations��

���ҵ��ǣ��Դ� UI ����̫����б���ȷ�Ĺ������Ҫ�������ظ���������롣

���ǿ���ͨ���Ƴ�һЩ������ִ�У�������һЩģ���Ƴ�ҵ���߼������磬ʹ�á����ﻯ����JavaScript �е� [`bind`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind)����Ȼ�����ǿ��ԴӺ��ĵĺ������洫�� state��������û����������ˡ�

����������û�м���������룬�����ٰ����ӹؼ�ҵ���߼��а��롣

```js
function FancyUserList(users) {
  return FancyBox(
    UserList.bind(null, users)
  );
}

const box = FancyUserList(data.users);
const resolvedChildren = box.children(likesPerUser, updateUserLikes);
const resolvedBox = {
  ...box,
  children: resolvedChildren
};
```

## State Map

֮ǰ����֪������ʹ����ϱ����ظ�ִ����ͬ�Ķ�������һ���ظ�ģʽ�����ǿ��԰�ִ�кʹ��� state �߼�Ų���������úܶ�ĵͲ㼶�ĺ�����ȥ��

```js
function FancyBoxWithState(
  children,
  stateMap,
  updateState
) {
  return FancyBox(
    children.map(child => child.continuation(
      stateMap.get(child.key),
      updateState
    ))
  );
}

function UserList(users) {
  return users.map(user => {
    continuation: FancyNameBox.bind(null, user),
    key: user.id
  });
}

function FancyUserList(users) {
  return FancyBoxWithState.bind(null,
    UserList(users)
  );
}

const continuation = FancyUserList(data.users);
continuation(likesPerUser, updateUserLikes);
```

## Memoization Map

һ����������һ�� memoization �б��� memoize ��� item �ͻ��ú����ѡ���Ϊ����Ҫ�ƶ����ӵĻ����㷨��ƽ�����Ƶ�ʺ��ڴ�ռ���ʡ�

���� UI ��ͬһ��λ�û���Ե��ȶ�����ͬ��λ��һ��ÿ�ζ��������ͬ�Ĳ���������������ʹ��һ���������� memoization ��һ���ǳ����õĲ��ԡ�

���ǿ����öԴ� state ͬ���ķ�ʽ������ϵĺ����д���һ�� memoization ���档

```js
function memoize(fn) {
  return function(arg, memoizationCache) {
    if (memoizationCache.arg === arg) {
      return memoizationCache.result;
    }
    const result = fn(arg);
    memoizationCache.arg = arg;
    memoizationCache.result = result;
    return result;
  };
}

function FancyBoxWithState(
  children,
  stateMap,
  updateState,
  memoizationCache
) {
  return FancyBox(
    children.map(child => child.continuation(
      stateMap.get(child.key),
      updateState,
      memoizationCache.get(child.key)
    ))
  );
}

const MemoizedFancyNameBox = memoize(FancyNameBox);
```

## ����ЧӦ��Algebraic Effects��

��������Ҫ������������ʱ��һ��㴫�����ݷǳ��鷳���������һ�ַ�ʽ�����ڶ������п�ݵش������ݣ�ͬʱ�ֲ���Ҫǣ�浽�м�㼶���Ǹ��ж�á�React �����ǰ���������context����

��ʱ�����������������ϸ��ճ��������϶��½��С��ٸ����ӣ��ڲ����㷨�У�����Ҫ��ʵ�����ǵ�λ��֮ǰ�˽��ӽڵ�Ĵ�С��

���ڣ����������һ�㳬�١��һ�ʹ�� [����ЧӦ](http://math.andrej.com/eff/) ������ҷ���� [ECMAScript ����������](https://esdiscuss.org/topic/one-shot-delimited-continuations-with-effect-handlers)�������Ժ���ʽ��̺���Ϥ������ �ڱ����� monad ǿ���������ʽһ���ı��롣

```js
function ThemeBorderColorRequest() { }

function FancyBox(children) {
  const color = raise new ThemeBorderColorRequest();
  return {
    borderWidth: '1px',
    borderColor: color,
    children: children
  };
}

function BlueTheme(children) {
  return try {
    children();
  } catch effect ThemeBorderColorRequest -> [, continuation] {
    continuation('blue');
  }
}

function App(data) {
  return BlueTheme(
    FancyUserList.bind(null, data.users)
  );
}
```