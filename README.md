# I See

Generate Ascii art from image.

## Installation

Require go 1.22+

```shell
go install github.com/mawngo/isee@latest
```

## Usage

convert image

```shell
> isee .\my-image.jpeg
```

or convert directory of images

```shell
> isee .\my-dir
```

### Options

```
> isee gray --help  
Generate grayscale ascii art

Usage:
  isee grayscale [flags]

Aliases:
  grayscale, gray

Flags:
  -h, --help                 help for grayscale
      --ramps string         Character ramps used to generate image
  -v, --vertical-ratio int   Shrinking height of the image (to % of the original size) (default 35)
  -w, --width uint           Width of the ascii art (default 80)

Global Flags:
      --debug   Enable debug mode

```

## Examples

```shell
> isee gray .\saturn.jpg
```

```
6:34PM INF Processing img=saturn.jpg dimension=240x240 width=80
                                                                                
                                                               .^;i+?}1)1{[_,   
                                                          `;<])/jrxxrjjuUzurt^  
                                                     ':<[|rvYJQOZwqpqwL({0Yu/`  
                                                 ';+)jzJ0mqppqqwwwqpdkow+0Yrl   
                                             .:+(xY0mqm0Uuf|1{1\cL0wqkafcLf;    
                                  .'^":;III<)nY0mmQzt}~l,"^^",:;-COqkbuzJ}`     
                             '"l~]{(\\\/jvJOww0Yx\[+l^      '":ljOqkOvUrl       
                          ^!-1/frrrjjnXQmqmCznxnvccur\[i^   `,ixmdqUXu+         
                       'l](frxxxrxvUOqqOJXYLmpdpwOQCYcnj(_"^;{CpqLXn-'          
                     'i{/jxnxxxvJZqqOJJ0d*MMohbpqwZQJXuxf\}~vpqQYx_'            
                    ,]\frrrxvUZpq0UUZhW8&M*oahbpmQJXcvunrf/{jQXt<.              
                   :{/ffjnY0qpZUzLdM8&WM#*ohdw0CUYXXzcvunrt(]{l                 
                  '](/jcCmpwCcXm*&WMWWM#obm0LJJJUUUYXXcvnrt(?.                  
                  ")rX0qp0znYdMW##MWW#kw0LCJJJJJJJJUYzcnrt|{<                   
                .<tYOpwJnrYp**aaooakqOLCJJCCCCCCCJUXcuxjt|{_`                   
              `?rU0pwXt\u0ddppdddw0LJJJJJCCCCLCJYzcuxf/(}?>'                    
            "}uUOpwx]]|nULLQOZZOCYYYYUUJJJCCJYzcvnj/({]_>,                      
          "{vY0pqz-,',-\xucXUUzcccczXXXXYYXcvunjt|1[?~!".                       
        `[czCppO/l"'   ^>[\jxjjrxnuuuuvvunxrjt|1[_<l,'                          
       ~vYcwbwQ(;:"'      ^!_{(|\\/ttt//\(){]_<I"`                              
     ,\LuXhbw0Y<;:,"^^";i_)rYOOCvf1>;;::,"`'.                                   
    >nQjcobqw0Cv\))(/rzCZwwZLzj{<,                                              
   ~uJJ+kakdppwwqppppwZQUvt}>,                                                  
  :jvUL]/LmwqwmZ0LUzuf)-!"                                                      
  ;fxuXUnjrxxxrf|{_i,'                                                          
   :+[}{{}]_<l,`                                                                
                                                                                


```