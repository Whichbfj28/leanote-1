# How to translate Leanote into other language?

Suppose that you want to add en-dk (the language is English and region is Danmark) to Leanote.

## 1. Copy `/messages/en-us` to `/messages/en-dk`
And translate all .conf files

## 2. Build to generate js i18n file
```
cd PATH-TO-LEANOTE
gulp
```