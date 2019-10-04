# Changelog

## 1.0.8 (09 Feb 2017)

### New
- option to specify an interval (in seconds) instead of numcaps (`--interval`)
- option which tries to skip ending credits by cuttin of 4 minutes of the movie (`--skip-credits`)

### Changes
- print errors when uploading fails
- with 1.0.8 the default behaviour of mt changed, `skip-credits` is now an opt-in and not the default anymore

## 1.0.7 (20 Nov 2016)

### New
- options for blur and blank threshold
- option to upload generated image to a given url (`--upload` and `--upload-url`)

### Changes
- improve usage of different skip functions when used in combination
- changed help message for some flags
- dont append `{{Count}}` to filename when useing `--single-images` with `--numcap 1`
- Fix an error where Resulution wasn't correctly added to header

## 1.0.6 (10 Jun 2016)

### New
- option to generate WEBVTT Files (`--webvtt`)

## 1.0.5 (07 Jun 2016)

### New
- option to put a watermark on the bottom-left corner of each image (`--watermark-all`)
- option to append a comment to the header area (`--comment`)
- option to list used config values (`--show-config`)
- experimental function for blur detection (`--skip-blurry`)

### Changes
- fixed a typo in config for `skip-existing` option
- compatible with go 1.6 and ffmpeg 3.0

## 1.0.4 (07 Jan 2016)

### New
- option for faster but more inaccurate seeking.
- experimental nude detection to skip images which are considered nude
- dont overwrite existing images by default (increments filename by 1 by default)
- switch for overwriting and skipping existing images

### Changes
- log used config file and values in debug mode
- update some of the dependencies
- use recent version of godeps to link to the saved deps

## 1.0.3 (24 Aug 2015)

### New
- option to save current settings to a settings file
- include default font and images in the binary
- added new filters
- added `--to` and `--from` options to use specific parts of the video file only
- option to provide a custom output path

## 1.0.2 (24 Aug 2015)

### New
- new filters: "fancy", "sephia" and "cross"
- option to include meta data in the header (bitrate, FPS and codecs)

### Change
- use all available cores to improve speed

## 1.0.1 (22 Jul 2015)

### Change
- hotfix to don't save images that are considered to black

## 1.0.0 (22 Jul 2015)

Inital release
