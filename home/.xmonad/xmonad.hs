import XMonad
import XMonad.Util.EZConfig

term = "gnome-terminal"
browser = "google-chrome"

main = xmonad $ defaultConfig {
    terminal    = term,
    borderWidth = 2,
    startupHook = startup
} `additionalKeysP` [ 
    ("C-<Return>",              spawn term),
    ("C-M-<Return>",            spawn browser),
    ("M-q",                     kill),
    ("M-S-r",                   restart "xmonad" True),
    ("<XF86AudioRaiseVolume>",  spawn "amixer -D pulse sset Master 5%+"),
    ("<XF86AudioLowerVolume>",  spawn "amixer -D pulse sset Master 5%-"),
    ("<XF86AudioMute>",         spawn "amixer -D pulse sset Master toggle"),
    ("M4-<F5>",   spawn "xbacklight -inc 10"),
    ("M4-<F6>",   spawn "xbacklight -dec 10")]


startup :: X ()
startup = do
    spawn "nm-applet"
    spawn "setxkbmap -option ctrl:nocaps"
    spawn "xcape -e 'Control_L=Escape"

