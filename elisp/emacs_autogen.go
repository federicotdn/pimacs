// Automatically generated with pimacs.extract
// DO NOT MODIFY!
// Generated from GNU Emacs commit: 3f43a16bc63eac12db5707b26da2507077b4f13c, branch: emacs-29

package elisp

const emacsConstant_coding_arg_max_autogen = 13
const emacsConstant_charset_arg_max_autogen = 17

func (ec *execContext) activeMinibufferWindow_autogen() (lispObject, error) {
	return ec.stub("active-minibuffer-window")
}

func (ec *execContext) setMinibufferWindow_autogen(window lispObject) (lispObject, error) {
	return ec.stub("set-minibuffer-window")
}

func (ec *execContext) minibufferp_autogen(buffer, live lispObject) (lispObject, error) {
	return ec.stub("minibufferp")
}

func (ec *execContext) innermostMinibufferP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("innermost-minibuffer-p")
}

func (ec *execContext) minibufferInnermostCommandLoopP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("minibuffer-innermost-command-loop-p")
}

func (ec *execContext) abortMinibuffers_autogen() (lispObject, error) {
	return ec.stub("abort-minibuffers")
}

func (ec *execContext) minibufferPromptEnd_autogen() (lispObject, error) {
	return ec.stub("minibuffer-prompt-end")
}

func (ec *execContext) minibufferContents_autogen() (lispObject, error) {
	return ec.stub("minibuffer-contents")
}

func (ec *execContext) minibufferContentsNoProperties_autogen() (lispObject, error) {
	return ec.stub("minibuffer-contents-no-properties")
}

func (ec *execContext) readFromMinibuffer_autogen(prompt, initial_contents, keymap, read, hist, default_value, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("read-from-minibuffer")
}

func (ec *execContext) readString_autogen(prompt, initial_input, history, default_value, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("read-string")
}

func (ec *execContext) readCommand_autogen(prompt, default_value lispObject) (lispObject, error) {
	return ec.stub("read-command")
}

func (ec *execContext) readFunction_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("read-function")
}

func (ec *execContext) readVariable_autogen(prompt, default_value lispObject) (lispObject, error) {
	return ec.stub("read-variable")
}

func (ec *execContext) readBuffer_autogen(prompt, def, require_match, predicate lispObject) (lispObject, error) {
	return ec.stub("read-buffer")
}

func (ec *execContext) tryCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return ec.stub("try-completion")
}

func (ec *execContext) allCompletions_autogen(string, collection, predicate, hide_spaces lispObject) (lispObject, error) {
	return ec.stub("all-completions")
}

func (ec *execContext) completingRead_autogen(prompt, collection, predicate, require_match, initial_input, hist, def, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("completing-read")
}

func (ec *execContext) testCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return ec.stub("test-completion")
}

func (ec *execContext) internalCompleteBuffer_autogen(string, predicate, flag lispObject) (lispObject, error) {
	return ec.stub("internal-complete-buffer")
}

func (ec *execContext) assocString_autogen(key, list, case_fold lispObject) (lispObject, error) {
	return ec.stub("assoc-string")
}

func (ec *execContext) minibufferDepth_autogen() (lispObject, error) {
	return ec.stub("minibuffer-depth")
}

func (ec *execContext) minibufferPrompt_autogen() (lispObject, error) {
	return ec.stub("minibuffer-prompt")
}

func (ec *execContext) timeAdd_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-add")
}

func (ec *execContext) timeSubtract_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-subtract")
}

func (ec *execContext) timeLessP_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-less-p")
}

func (ec *execContext) timeEqualP_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-equal-p")
}

func (ec *execContext) floatTime_autogen(specified_time lispObject) (lispObject, error) {
	return ec.stub("float-time")
}

func (ec *execContext) formatTimeString_autogen(format_string, timeval, zone lispObject) (lispObject, error) {
	return ec.stub("format-time-string")
}

func (ec *execContext) decodeTime_autogen(specified_time, zone, form lispObject) (lispObject, error) {
	return ec.stub("decode-time")
}

func (ec *execContext) encodeTime_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("encode-time")
}

func (ec *execContext) timeConvert_autogen(time, form lispObject) (lispObject, error) {
	return ec.stub("time-convert")
}

func (ec *execContext) currentTime_autogen() (lispObject, error) {
	return ec.stub("current-time")
}

func (ec *execContext) currentCpuTime_autogen() (lispObject, error) {
	return ec.stub("current-cpu-time")
}

func (ec *execContext) currentTimeString_autogen(specified_time, zone lispObject) (lispObject, error) {
	return ec.stub("current-time-string")
}

func (ec *execContext) currentTimeZone_autogen(specified_time, zone lispObject) (lispObject, error) {
	return ec.stub("current-time-zone")
}

func (ec *execContext) setTimeZoneRule_autogen(tz lispObject) (lispObject, error) {
	return ec.stub("set-time-zone-rule")
}

func (ec *execContext) ttyDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-display-color-p")
}

func (ec *execContext) ttyDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-display-color-cells")
}

func (ec *execContext) ttyType_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-type")
}

func (ec *execContext) controllingTtyP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("controlling-tty-p")
}

func (ec *execContext) ttyNoUnderline_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-no-underline")
}

func (ec *execContext) ttyTopFrame_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-top-frame")
}

func (ec *execContext) suspendTty_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("suspend-tty")
}

func (ec *execContext) resumeTty_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("resume-tty")
}

func (ec *execContext) ttySetOutputBufferSize_autogen(size, tty lispObject) (lispObject, error) {
	return ec.stub("tty--set-output-buffer-size")
}

func (ec *execContext) ttyOutputBufferSize_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("tty--output-buffer-size")
}

func (ec *execContext) gpmMouseStart_autogen() (lispObject, error) {
	return ec.stub("gpm-mouse-start")
}

func (ec *execContext) gpmMouseStop_autogen() (lispObject, error) {
	return ec.stub("gpm-mouse-stop")
}

func (ec *execContext) undoBoundary_autogen() (lispObject, error) {
	return ec.stub("undo-boundary")
}

func (ec *execContext) startKbdMacro_autogen(append, no_exec lispObject) (lispObject, error) {
	return ec.stub("start-kbd-macro")
}

func (ec *execContext) endKbdMacro_autogen(repeat, loopfunc lispObject) (lispObject, error) {
	return ec.stub("end-kbd-macro")
}

func (ec *execContext) cancelKbdMacroEvents_autogen() (lispObject, error) {
	return ec.stub("cancel-kbd-macro-events")
}

func (ec *execContext) storeKbdMacroEvent_autogen(event lispObject) (lispObject, error) {
	return ec.stub("store-kbd-macro-event")
}

func (ec *execContext) callLastKbdMacro_autogen(prefix, loopfunc lispObject) (lispObject, error) {
	return ec.stub("call-last-kbd-macro")
}

func (ec *execContext) executeKbdMacro_autogen(macro, count, loopfunc lispObject) (lispObject, error) {
	return ec.stub("execute-kbd-macro")
}

func (ec *execContext) libxmlParseHtmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return ec.stub("libxml-parse-html-region")
}

func (ec *execContext) libxmlParseXmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return ec.stub("libxml-parse-xml-region")
}

func (ec *execContext) libxmlAvailableP_autogen() (lispObject, error) {
	return ec.stub("libxml-available-p")
}

func (ec *execContext) kqueueAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return ec.stub("kqueue-add-watch")
}

func (ec *execContext) kqueueRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("kqueue-rm-watch")
}

func (ec *execContext) kqueueValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("kqueue-valid-p")
}

func (ec *execContext) combineAfterChangeExecute_autogen() (lispObject, error) {
	return ec.stub("combine-after-change-execute")
}

func (ec *execContext) cclProgramP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("ccl-program-p")
}

func (ec *execContext) cclExecute_autogen(ccl_prog, reg lispObject) (lispObject, error) {
	return ec.stub("ccl-execute")
}

func (ec *execContext) cclExecuteOnString_autogen(ccl_prog, status, str, contin, unibyte_p lispObject) (lispObject, error) {
	return ec.stub("ccl-execute-on-string")
}

func (ec *execContext) registerCclProgram_autogen(name, ccl_prog lispObject) (lispObject, error) {
	return ec.stub("register-ccl-program")
}

func (ec *execContext) registerCodeConversionMap_autogen(symbol, map_ lispObject) (lispObject, error) {
	return ec.stub("register-code-conversion-map")
}

func (ec *execContext) syntaxTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("syntax-table-p")
}

func (ec *execContext) syntaxTable_autogen() (lispObject, error) {
	return ec.stub("syntax-table")
}

func (ec *execContext) standardSyntaxTable_autogen() (lispObject, error) {
	return ec.stub("standard-syntax-table")
}

func (ec *execContext) copySyntaxTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-syntax-table")
}

func (ec *execContext) setSyntaxTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-syntax-table")
}

func (ec *execContext) charSyntax_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-syntax")
}

func (ec *execContext) syntaxClassToChar_autogen(syntax lispObject) (lispObject, error) {
	return ec.stub("syntax-class-to-char")
}

func (ec *execContext) matchingParen_autogen(character lispObject) (lispObject, error) {
	return ec.stub("matching-paren")
}

func (ec *execContext) stringToSyntax_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-syntax")
}

func (ec *execContext) modifySyntaxEntry_autogen(c, newentry, syntax_table lispObject) (lispObject, error) {
	return ec.stub("modify-syntax-entry")
}

func (ec *execContext) internalDescribeSyntaxValue_autogen(syntax lispObject) (lispObject, error) {
	return ec.stub("internal-describe-syntax-value")
}

func (ec *execContext) forwardWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("forward-word")
}

func (ec *execContext) skipCharsForward_autogen(string, lim lispObject) (lispObject, error) {
	return ec.stub("skip-chars-forward")
}

func (ec *execContext) skipCharsBackward_autogen(string, lim lispObject) (lispObject, error) {
	return ec.stub("skip-chars-backward")
}

func (ec *execContext) skipSyntaxForward_autogen(syntax, lim lispObject) (lispObject, error) {
	return ec.stub("skip-syntax-forward")
}

func (ec *execContext) skipSyntaxBackward_autogen(syntax, lim lispObject) (lispObject, error) {
	return ec.stub("skip-syntax-backward")
}

func (ec *execContext) forwardComment_autogen(count lispObject) (lispObject, error) {
	return ec.stub("forward-comment")
}

func (ec *execContext) scanLists_autogen(from, count, depth lispObject) (lispObject, error) {
	return ec.stub("scan-lists")
}

func (ec *execContext) scanSexps_autogen(from, count lispObject) (lispObject, error) {
	return ec.stub("scan-sexps")
}

func (ec *execContext) backwardPrefixChars_autogen() (lispObject, error) {
	return ec.stub("backward-prefix-chars")
}

func (ec *execContext) parsePartialSexp_autogen(from, to, targetdepth, stopbefore, oldstate, commentstop lispObject) (lispObject, error) {
	return ec.stub("parse-partial-sexp")
}

func (ec *execContext) clearCompositionCache_autogen() (lispObject, error) {
	return ec.stub("clear-composition-cache")
}

func (ec *execContext) compositionGetGstring_autogen(from, to, font_object, string lispObject) (lispObject, error) {
	return ec.stub("composition-get-gstring")
}

func (ec *execContext) composeRegionInternal_autogen(start, end, components, modification_func lispObject) (lispObject, error) {
	return ec.stub("compose-region-internal")
}

func (ec *execContext) composeStringInternal_autogen(string, start, end, components, modification_func lispObject) (lispObject, error) {
	return ec.stub("compose-string-internal")
}

func (ec *execContext) findCompositionInternal_autogen(pos, limit, string, detail_p lispObject) (lispObject, error) {
	return ec.stub("find-composition-internal")
}

func (ec *execContext) compositionSortRules_autogen(rules lispObject) (lispObject, error) {
	return ec.stub("composition-sort-rules")
}

func (ec *execContext) inotifyAddWatch_autogen(filename, aspect, callback lispObject) (lispObject, error) {
	return ec.stub("inotify-add-watch")
}

func (ec *execContext) inotifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("inotify-rm-watch")
}

func (ec *execContext) inotifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("inotify-valid-p")
}

func (ec *execContext) inotifyWatchList_autogen() (lispObject, error) {
	return ec.stub("inotify-watch-list")
}

func (ec *execContext) inotifyAllocatedP_autogen() (lispObject, error) {
	return ec.stub("inotify-allocated-p")
}

func (ec *execContext) charToString_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-to-string")
}

func (ec *execContext) byteToString_autogen(byte lispObject) (lispObject, error) {
	return ec.stub("byte-to-string")
}

func (ec *execContext) stringToChar_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-char")
}

func (ec *execContext) point_autogen() (lispObject, error) {
	return ec.stub("point")
}

func (ec *execContext) pointMarker_autogen() (lispObject, error) {
	return ec.stub("point-marker")
}

func (ec *execContext) gotoChar_autogen(position lispObject) (lispObject, error) {
	return ec.stub("goto-char")
}

func (ec *execContext) regionBeginning_autogen() (lispObject, error) {
	return ec.stub("region-beginning")
}

func (ec *execContext) regionEnd_autogen() (lispObject, error) {
	return ec.stub("region-end")
}

func (ec *execContext) markMarker_autogen() (lispObject, error) {
	return ec.stub("mark-marker")
}

func (ec *execContext) getPosProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-pos-property")
}

func (ec *execContext) deleteField_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("delete-field")
}

func (ec *execContext) fieldString_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("field-string")
}

func (ec *execContext) fieldStringNoProperties_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("field-string-no-properties")
}

func (ec *execContext) fieldBeginning_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return ec.stub("field-beginning")
}

func (ec *execContext) fieldEnd_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return ec.stub("field-end")
}

func (ec *execContext) constrainToField_autogen(new_pos, old_pos, escape_from_edge, only_in_line, inhibit_capture_property lispObject) (lispObject, error) {
	return ec.stub("constrain-to-field")
}

func (ec *execContext) posBol_autogen(n lispObject) (lispObject, error) {
	return ec.stub("pos-bol")
}

func (ec *execContext) lineBeginningPosition_autogen(n lispObject) (lispObject, error) {
	return ec.stub("line-beginning-position")
}

func (ec *execContext) posEol_autogen(n lispObject) (lispObject, error) {
	return ec.stub("pos-eol")
}

func (ec *execContext) lineEndPosition_autogen(n lispObject) (lispObject, error) {
	return ec.stub("line-end-position")
}

func (ec *execContext) saveExcursion_autogen(args lispObject) (lispObject, error) {
	return ec.stub("save-excursion")
}

func (ec *execContext) saveCurrentBuffer_autogen(args lispObject) (lispObject, error) {
	return ec.stub("save-current-buffer")
}

func (ec *execContext) bufferSize_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-size")
}

func (ec *execContext) pointMin_autogen() (lispObject, error) {
	return ec.stub("point-min")
}

func (ec *execContext) pointMinMarker_autogen() (lispObject, error) {
	return ec.stub("point-min-marker")
}

func (ec *execContext) pointMax_autogen() (lispObject, error) {
	return ec.stub("point-max")
}

func (ec *execContext) pointMaxMarker_autogen() (lispObject, error) {
	return ec.stub("point-max-marker")
}

func (ec *execContext) gapPosition_autogen() (lispObject, error) {
	return ec.stub("gap-position")
}

func (ec *execContext) gapSize_autogen() (lispObject, error) {
	return ec.stub("gap-size")
}

func (ec *execContext) positionBytes_autogen(position lispObject) (lispObject, error) {
	return ec.stub("position-bytes")
}

func (ec *execContext) byteToPosition_autogen(bytepos lispObject) (lispObject, error) {
	return ec.stub("byte-to-position")
}

func (ec *execContext) followingChar_autogen() (lispObject, error) {
	return ec.stub("following-char")
}

func (ec *execContext) previousChar_autogen() (lispObject, error) {
	return ec.stub("preceding-char")
}

func (ec *execContext) bobp_autogen() (lispObject, error) {
	return ec.stub("bobp")
}

func (ec *execContext) eobp_autogen() (lispObject, error) {
	return ec.stub("eobp")
}

func (ec *execContext) bolp_autogen() (lispObject, error) {
	return ec.stub("bolp")
}

func (ec *execContext) eolp_autogen() (lispObject, error) {
	return ec.stub("eolp")
}

func (ec *execContext) charAfter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("char-after")
}

func (ec *execContext) charBefore_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("char-before")
}

func (ec *execContext) userLoginName_autogen(uid lispObject) (lispObject, error) {
	return ec.stub("user-login-name")
}

func (ec *execContext) userRealLoginName_autogen() (lispObject, error) {
	return ec.stub("user-real-login-name")
}

func (ec *execContext) userUid_autogen() (lispObject, error) {
	return ec.stub("user-uid")
}

func (ec *execContext) userRealUid_autogen() (lispObject, error) {
	return ec.stub("user-real-uid")
}

func (ec *execContext) groupName_autogen(gid lispObject) (lispObject, error) {
	return ec.stub("group-name")
}

func (ec *execContext) groupGid_autogen() (lispObject, error) {
	return ec.stub("group-gid")
}

func (ec *execContext) groupRealGid_autogen() (lispObject, error) {
	return ec.stub("group-real-gid")
}

func (ec *execContext) userFullName_autogen(uid lispObject) (lispObject, error) {
	return ec.stub("user-full-name")
}

func (ec *execContext) systemName_autogen() (lispObject, error) {
	return ec.stub("system-name")
}

func (ec *execContext) emacsPid_autogen() (lispObject, error) {
	return ec.stub("emacs-pid")
}

func (ec *execContext) insert_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert")
}

func (ec *execContext) insertAndInherit_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-and-inherit")
}

func (ec *execContext) insertBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-before-markers")
}

func (ec *execContext) insertAndInheritBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-before-markers-and-inherit")
}

func (ec *execContext) insertChar_autogen(character, count, inherit lispObject) (lispObject, error) {
	return ec.stub("insert-char")
}

func (ec *execContext) insertByte_autogen(byte, count, inherit lispObject) (lispObject, error) {
	return ec.stub("insert-byte")
}

func (ec *execContext) bufferSubstring_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("buffer-substring")
}

func (ec *execContext) bufferSubstringNoProperties_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("buffer-substring-no-properties")
}

func (ec *execContext) bufferString_autogen() (lispObject, error) {
	return ec.stub("buffer-string")
}

func (ec *execContext) insertBufferSubstring_autogen(buffer, start, end lispObject) (lispObject, error) {
	return ec.stub("insert-buffer-substring")
}

func (ec *execContext) compareBufferSubstrings_autogen(buffer1, start1, end1, buffer2, start2, end2 lispObject) (lispObject, error) {
	return ec.stub("compare-buffer-substrings")
}

func (ec *execContext) replaceBufferContents_autogen(source, max_secs, max_costs lispObject) (lispObject, error) {
	return ec.stub("replace-buffer-contents")
}

func (ec *execContext) substCharInRegion_autogen(start, end, fromchar, tochar, noundo lispObject) (lispObject, error) {
	return ec.stub("subst-char-in-region")
}

func (ec *execContext) translateRegionInternal_autogen(start, end, table lispObject) (lispObject, error) {
	return ec.stub("translate-region-internal")
}

func (ec *execContext) deleteRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("delete-region")
}

func (ec *execContext) deleteAndExtractRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("delete-and-extract-region")
}

func (ec *execContext) widen_autogen() (lispObject, error) {
	return ec.stub("widen")
}

func (ec *execContext) narrowToRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("narrow-to-region")
}

func (ec *execContext) internalLockNarrowing_autogen(tag lispObject) (lispObject, error) {
	return ec.stub("internal--lock-narrowing")
}

func (ec *execContext) internalUnlockNarrowing_autogen(tag lispObject) (lispObject, error) {
	return ec.stub("internal--unlock-narrowing")
}

func (ec *execContext) saveRestriction_autogen(body lispObject) (lispObject, error) {
	return ec.stub("save-restriction")
}

func (ec *execContext) ngettext_autogen(msgid, msgid_plural, n lispObject) (lispObject, error) {
	return ec.stub("ngettext")
}

func (ec *execContext) message_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message")
}

func (ec *execContext) messageBox_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message-box")
}

func (ec *execContext) messageOrBox_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message-or-box")
}

func (ec *execContext) currentMessage_autogen() (lispObject, error) {
	return ec.stub("current-message")
}

func (ec *execContext) propertize_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("propertize")
}

func (ec *execContext) format_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("format")
}

func (ec *execContext) formatMessage_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("format-message")
}

func (ec *execContext) charEqual_autogen(c1, c2 lispObject) (lispObject, error) {
	return ec.stub("char-equal")
}

func (ec *execContext) transposeRegions_autogen(startr1, endr1, startr2, endr2, leave_markers lispObject) (lispObject, error) {
	return ec.stub("transpose-regions")
}

func (ec *execContext) cygwinConvertFileNameToWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return ec.stub("cygwin-convert-file-name-to-windows")
}

func (ec *execContext) cygwinConvertFileNameFromWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return ec.stub("cygwin-convert-file-name-from-windows")
}

func (ec *execContext) recursiveEdit_autogen() (lispObject, error) {
	return ec.stub("recursive-edit")
}

func (ec *execContext) commandErrorDefaultFunction_autogen(data, context, signal lispObject) (lispObject, error) {
	return ec.stub("command-error-default-function")
}

func (ec *execContext) topLevel_autogen() (lispObject, error) {
	return ec.stub("top-level")
}

func (ec *execContext) exitRecursiveEdit_autogen() (lispObject, error) {
	return ec.stub("exit-recursive-edit")
}

func (ec *execContext) abortRecursiveEdit_autogen() (lispObject, error) {
	return ec.stub("abort-recursive-edit")
}

func (ec *execContext) internalTrackMouse_autogen(bodyfun lispObject) (lispObject, error) {
	return ec.stub("internal--track-mouse")
}

func (ec *execContext) currentIdleTime_autogen() (lispObject, error) {
	return ec.stub("current-idle-time")
}

func (ec *execContext) eventSymbolParseModifiers_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("internal-event-symbol-parse-modifiers")
}

func (ec *execContext) eventConvertList_autogen(event_desc lispObject) (lispObject, error) {
	return ec.stub("event-convert-list")
}

func (ec *execContext) internalHandleFocusIn_autogen(event lispObject) (lispObject, error) {
	return ec.stub("internal-handle-focus-in")
}

func (ec *execContext) readKeySequence_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop lispObject) (lispObject, error) {
	return ec.stub("read-key-sequence")
}

func (ec *execContext) readKeySequenceVector_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop lispObject) (lispObject, error) {
	return ec.stub("read-key-sequence-vector")
}

func (ec *execContext) inputPendingP_autogen(check_timers lispObject) (lispObject, error) {
	return ec.stub("input-pending-p")
}

func (ec *execContext) lossageSize_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("lossage-size")
}

func (ec *execContext) recentKeys_autogen(include_cmds lispObject) (lispObject, error) {
	return ec.stub("recent-keys")
}

func (ec *execContext) thisCommandKeys_autogen() (lispObject, error) {
	return ec.stub("this-command-keys")
}

func (ec *execContext) setThisCommandKeys_autogen(keys lispObject) (lispObject, error) {
	return ec.stub("set--this-command-keys")
}

func (ec *execContext) thisCommandKeysVector_autogen() (lispObject, error) {
	return ec.stub("this-command-keys-vector")
}

func (ec *execContext) thisSingleCommandKeys_autogen() (lispObject, error) {
	return ec.stub("this-single-command-keys")
}

func (ec *execContext) thisSingleCommandRawKeys_autogen() (lispObject, error) {
	return ec.stub("this-single-command-raw-keys")
}

func (ec *execContext) clearThisCommandKeys_autogen(keep_record lispObject) (lispObject, error) {
	return ec.stub("clear-this-command-keys")
}

func (ec *execContext) recursionDepth_autogen() (lispObject, error) {
	return ec.stub("recursion-depth")
}

func (ec *execContext) openDribbleFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("open-dribble-file")
}

func (ec *execContext) discardInput_autogen() (lispObject, error) {
	return ec.stub("discard-input")
}

func (ec *execContext) suspendEmacs_autogen(stuffstring lispObject) (lispObject, error) {
	return ec.stub("suspend-emacs")
}

func (ec *execContext) setInputInterruptMode_autogen(interrupt lispObject) (lispObject, error) {
	return ec.stub("set-input-interrupt-mode")
}

func (ec *execContext) setOutputFlowControl_autogen(flow, terminal lispObject) (lispObject, error) {
	return ec.stub("set-output-flow-control")
}

func (ec *execContext) setInputMetaMode_autogen(meta, terminal lispObject) (lispObject, error) {
	return ec.stub("set-input-meta-mode")
}

func (ec *execContext) setQuitChar_autogen(quit lispObject) (lispObject, error) {
	return ec.stub("set-quit-char")
}

func (ec *execContext) setInputMode_autogen(interrupt, flow, meta, quit lispObject) (lispObject, error) {
	return ec.stub("set-input-mode")
}

func (ec *execContext) currentInputMode_autogen() (lispObject, error) {
	return ec.stub("current-input-mode")
}

func (ec *execContext) posnAtXY_autogen(x, y, frame_or_window, whole lispObject) (lispObject, error) {
	return ec.stub("posn-at-x-y")
}

func (ec *execContext) posnAtPoint_autogen(pos, window lispObject) (lispObject, error) {
	return ec.stub("posn-at-point")
}

func (ec *execContext) sqliteOpen_autogen(file lispObject) (lispObject, error) {
	return ec.stub("sqlite-open")
}

func (ec *execContext) sqliteClose_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-close")
}

func (ec *execContext) sqliteExecute_autogen(db, query, values lispObject) (lispObject, error) {
	return ec.stub("sqlite-execute")
}

func (ec *execContext) sqliteSelect_autogen(db, query, values, return_type lispObject) (lispObject, error) {
	return ec.stub("sqlite-select")
}

func (ec *execContext) sqliteTransaction_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-transaction")
}

func (ec *execContext) sqliteCommit_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-commit")
}

func (ec *execContext) sqliteRollback_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-rollback")
}

func (ec *execContext) sqlitePragma_autogen(db, pragma lispObject) (lispObject, error) {
	return ec.stub("sqlite-pragma")
}

func (ec *execContext) sqliteLoadExtension_autogen(db, module lispObject) (lispObject, error) {
	return ec.stub("sqlite-load-extension")
}

func (ec *execContext) sqliteNext_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-next")
}

func (ec *execContext) sqliteColumns_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-columns")
}

func (ec *execContext) sqliteMoreP_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-more-p")
}

func (ec *execContext) sqliteFinalize_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-finalize")
}

func (ec *execContext) sqliteVersion_autogen() (lispObject, error) {
	return ec.stub("sqlite-version")
}

func (ec *execContext) sqlitep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("sqlitep")
}

func (ec *execContext) sqliteAvailableP_autogen() (lispObject, error) {
	return ec.stub("sqlite-available-p")
}

func (ec *execContext) menuOrPopupActiveP_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p")
}

func (ec *execContext) menuOrPopupActiveP_1_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p")
}

func (ec *execContext) menuOrPopupActiveP_2_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p")
}

func (ec *execContext) menuOrPopupActiveP_3_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p")
}

func (ec *execContext) haikuMenuBarOpen_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("haiku-menu-bar-open")
}

func (ec *execContext) writeChar_autogen(character, printcharfun lispObject) (lispObject, error) {
	return ec.stub("write-char")
}

func (ec *execContext) terpri_autogen(printcharfun, ensure lispObject) (lispObject, error) {
	return ec.stub("terpri")
}

func (ec *execContext) prin1_autogen(object, printcharfun, overrides lispObject) (lispObject, error) {
	return ec.stub("prin1")
}

func (ec *execContext) prin1ToString_autogen(object, noescape, overrides lispObject) (lispObject, error) {
	return ec.stub("prin1-to-string")
}

func (ec *execContext) princ_autogen(object, printcharfun lispObject) (lispObject, error) {
	return ec.stub("princ")
}

func (ec *execContext) print_autogen(object, printcharfun lispObject) (lispObject, error) {
	return ec.stub("print")
}

func (ec *execContext) flushStandardOutput_autogen() (lispObject, error) {
	return ec.stub("flush-standard-output")
}

func (ec *execContext) externalDebuggingOutput_autogen(character lispObject) (lispObject, error) {
	return ec.stub("external-debugging-output")
}

func (ec *execContext) redirectDebuggingOutput_autogen(file, append lispObject) (lispObject, error) {
	return ec.stub("redirect-debugging-output")
}

func (ec *execContext) errorMessageString_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("error-message-string")
}

func (ec *execContext) printPreprocess_autogen(object lispObject) (lispObject, error) {
	return ec.stub("print--preprocess")
}

func (ec *execContext) or_autogen(args lispObject) (lispObject, error) {
	return ec.stub("or")
}

func (ec *execContext) and_autogen(args lispObject) (lispObject, error) {
	return ec.stub("and")
}

func (ec *execContext) if_autogen(args lispObject) (lispObject, error) {
	return ec.stub("if")
}

func (ec *execContext) cond_autogen(args lispObject) (lispObject, error) {
	return ec.stub("cond")
}

func (ec *execContext) progn_autogen(body lispObject) (lispObject, error) {
	return ec.stub("progn")
}

func (ec *execContext) prog1_autogen(args lispObject) (lispObject, error) {
	return ec.stub("prog1")
}

func (ec *execContext) setq_autogen(args lispObject) (lispObject, error) {
	return ec.stub("setq")
}

func (ec *execContext) quote_autogen(args lispObject) (lispObject, error) {
	return ec.stub("quote")
}

func (ec *execContext) function_autogen(args lispObject) (lispObject, error) {
	return ec.stub("function")
}

func (ec *execContext) defvaralias_autogen(new_alias, base_variable, docstring lispObject) (lispObject, error) {
	return ec.stub("defvaralias")
}

func (ec *execContext) defaultToplevelValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-toplevel-value")
}

func (ec *execContext) setDefaultToplevelValue_autogen(symbol, value lispObject) (lispObject, error) {
	return ec.stub("set-default-toplevel-value")
}

func (ec *execContext) internalDefineUninitializedVariable_autogen(symbol, doc lispObject) (lispObject, error) {
	return ec.stub("internal--define-uninitialized-variable")
}

func (ec *execContext) defvar_autogen(args lispObject) (lispObject, error) {
	return ec.stub("defvar")
}

func (ec *execContext) defvar1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return ec.stub("defvar-1")
}

func (ec *execContext) defconst_autogen(args lispObject) (lispObject, error) {
	return ec.stub("defconst")
}

func (ec *execContext) defconst1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return ec.stub("defconst-1")
}

func (ec *execContext) makeVarNonSpecial_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("internal-make-var-non-special")
}

func (ec *execContext) letX_autogen(args lispObject) (lispObject, error) {
	return ec.stub("let*")
}

func (ec *execContext) let_autogen(args lispObject) (lispObject, error) {
	return ec.stub("let")
}

func (ec *execContext) while_autogen(args lispObject) (lispObject, error) {
	return ec.stub("while")
}

func (ec *execContext) funcallWithDelayedMessage_autogen(timeout, message, function lispObject) (lispObject, error) {
	return ec.stub("funcall-with-delayed-message")
}

func (ec *execContext) macroexpand_autogen(form, environment lispObject) (lispObject, error) {
	return ec.stub("macroexpand")
}

func (ec *execContext) catch_autogen(args lispObject) (lispObject, error) {
	return ec.stub("catch")
}

func (ec *execContext) throw_autogen(tag, value lispObject) (lispObject, error) {
	return ec.stub("throw")
}

func (ec *execContext) unwindProtect_autogen(args lispObject) (lispObject, error) {
	return ec.stub("unwind-protect")
}

func (ec *execContext) conditionCase_autogen(args lispObject) (lispObject, error) {
	return ec.stub("condition-case")
}

func (ec *execContext) signal_autogen(error_symbol, data lispObject) (lispObject, error) {
	return ec.stub("signal")
}

func (ec *execContext) commandp_autogen(function, for_call_interactively lispObject) (lispObject, error) {
	return ec.stub("commandp")
}

func (ec *execContext) autoload_autogen(function, file, docstring, interactive, type_ lispObject) (lispObject, error) {
	return ec.stub("autoload")
}

func (ec *execContext) autoloadDoLoad_autogen(fundef, funname, macro_only lispObject) (lispObject, error) {
	return ec.stub("autoload-do-load")
}

func (ec *execContext) eval_autogen(form, lexical lispObject) (lispObject, error) {
	return ec.stub("eval")
}

func (ec *execContext) apply_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("apply")
}

func (ec *execContext) runHooks_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hooks")
}

func (ec *execContext) runHookWithArgs_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args")
}

func (ec *execContext) runHookWithArgsUntilSuccess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args-until-success")
}

func (ec *execContext) runHookWithArgsUntilFailure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args-until-failure")
}

func (ec *execContext) runHookWrapped_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-wrapped")
}

func (ec *execContext) functionp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("functionp")
}

func (ec *execContext) funcall_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("funcall")
}

func (ec *execContext) funcArity_autogen(function lispObject) (lispObject, error) {
	return ec.stub("func-arity")
}

func (ec *execContext) fetchBytecode_autogen(object lispObject) (lispObject, error) {
	return ec.stub("fetch-bytecode")
}

func (ec *execContext) specialVariableP_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("special-variable-p")
}

func (ec *execContext) backtraceDebug_autogen(level, flag lispObject) (lispObject, error) {
	return ec.stub("backtrace-debug")
}

func (ec *execContext) mapbacktrace_autogen(function, base lispObject) (lispObject, error) {
	return ec.stub("mapbacktrace")
}

func (ec *execContext) backtraceFrameInternal_autogen(function, nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace-frame--internal")
}

func (ec *execContext) backtraceFramesFromThread_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("backtrace--frames-from-thread")
}

func (ec *execContext) backtraceEval_autogen(exp, nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace-eval")
}

func (ec *execContext) backtraceLocals_autogen(nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace--locals")
}

func (ec *execContext) xMenuBarOpenInternal_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal")
}

func (ec *execContext) xMenuBarOpenInternal_1_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal")
}

func (ec *execContext) xMenuBarOpenInternal_2_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal")
}

func (ec *execContext) int86_autogen(interrupt, registers lispObject) (lispObject, error) {
	return ec.stub("int86")
}

func (ec *execContext) dosMemget_autogen(address, vector lispObject) (lispObject, error) {
	return ec.stub("msdos-memget")
}

func (ec *execContext) dosMemput_autogen(address, vector lispObject) (lispObject, error) {
	return ec.stub("msdos-memput")
}

func (ec *execContext) msdosSetKeyboard_autogen(country_code, allkeys lispObject) (lispObject, error) {
	return ec.stub("msdos-set-keyboard")
}

func (ec *execContext) msdosMouseP_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-p")
}

func (ec *execContext) msdosMouseInit_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-init")
}

func (ec *execContext) msdosMouseEnable_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-enable")
}

func (ec *execContext) msdosMouseDisable_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-disable")
}

func (ec *execContext) insertStartupScreen_autogen() (lispObject, error) {
	return ec.stub("insert-startup-screen")
}

func (ec *execContext) fileSystemInfo_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info")
}

func (ec *execContext) fileSystemInfo_1_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info")
}

func (ec *execContext) fileSystemInfo_2_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info")
}

func (ec *execContext) textPropertiesAt_autogen(position, object lispObject) (lispObject, error) {
	return ec.stub("text-properties-at")
}

func (ec *execContext) getTextProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-text-property")
}

func (ec *execContext) getCharProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-char-property")
}

func (ec *execContext) getCharPropertyAndOverlay_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-char-property-and-overlay")
}

func (ec *execContext) nextCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return ec.stub("next-char-property-change")
}

func (ec *execContext) previousCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return ec.stub("previous-char-property-change")
}

func (ec *execContext) nextSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-single-char-property-change")
}

func (ec *execContext) previousSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-single-char-property-change")
}

func (ec *execContext) nextPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-property-change")
}

func (ec *execContext) nextSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-single-property-change")
}

func (ec *execContext) previousPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-property-change")
}

func (ec *execContext) previousSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-single-property-change")
}

func (ec *execContext) addTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("add-text-properties")
}

func (ec *execContext) putTextProperty_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("put-text-property")
}

func (ec *execContext) setTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("set-text-properties")
}

func (ec *execContext) addFaceTextProperty_autogen(start, end, face, append, object lispObject) (lispObject, error) {
	return ec.stub("add-face-text-property")
}

func (ec *execContext) removeTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("remove-text-properties")
}

func (ec *execContext) removeListOfTextProperties_autogen(start, end, list_of_properties, object lispObject) (lispObject, error) {
	return ec.stub("remove-list-of-text-properties")
}

func (ec *execContext) textPropertyAny_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("text-property-any")
}

func (ec *execContext) textPropertyNotAll_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("text-property-not-all")
}

func (ec *execContext) makeMutex_autogen(name lispObject) (lispObject, error) {
	return ec.stub("make-mutex")
}

func (ec *execContext) mutexLock_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-lock")
}

func (ec *execContext) mutexUnlock_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-unlock")
}

func (ec *execContext) mutexName_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-name")
}

func (ec *execContext) makeConditionVariable_autogen(mutex, name lispObject) (lispObject, error) {
	return ec.stub("make-condition-variable")
}

func (ec *execContext) conditionWait_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-wait")
}

func (ec *execContext) conditionNotify_autogen(cond, all lispObject) (lispObject, error) {
	return ec.stub("condition-notify")
}

func (ec *execContext) conditionMutex_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-mutex")
}

func (ec *execContext) conditionName_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-name")
}

func (ec *execContext) threadYield_autogen() (lispObject, error) {
	return ec.stub("thread-yield")
}

func (ec *execContext) makeThread_autogen(function, name lispObject) (lispObject, error) {
	return ec.stub("make-thread")
}

func (ec *execContext) currentThread_autogen() (lispObject, error) {
	return ec.stub("current-thread")
}

func (ec *execContext) threadName_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-name")
}

func (ec *execContext) threadSignal_autogen(thread, error_symbol, data lispObject) (lispObject, error) {
	return ec.stub("thread-signal")
}

func (ec *execContext) threadLiveP_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-live-p")
}

func (ec *execContext) threadBlocker_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread--blocker")
}

func (ec *execContext) threadJoin_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-join")
}

func (ec *execContext) allThreads_autogen() (lispObject, error) {
	return ec.stub("all-threads")
}

func (ec *execContext) threadLastError_autogen(cleanup lispObject) (lispObject, error) {
	return ec.stub("thread-last-error")
}

func (ec *execContext) fontGetSystemNormalFont_autogen() (lispObject, error) {
	return ec.stub("font-get-system-normal-font")
}

func (ec *execContext) fontGetSystemNormalFont_1_autogen() (lispObject, error) {
	return ec.stub("font-get-system-normal-font")
}

func (ec *execContext) fontGetSystemFont_autogen() (lispObject, error) {
	return ec.stub("font-get-system-font")
}

func (ec *execContext) fontGetSystemFont_1_autogen() (lispObject, error) {
	return ec.stub("font-get-system-font")
}

func (ec *execContext) toolBarGetSystemStyle_autogen() (lispObject, error) {
	return ec.stub("tool-bar-get-system-style")
}

func (ec *execContext) makeCharTable_autogen(purpose, init lispObject) (lispObject, error) {
	return ec.stub("make-char-table")
}

func (ec *execContext) charTableSubtype_autogen(char_table lispObject) (lispObject, error) {
	return ec.stub("char-table-subtype")
}

func (ec *execContext) charTableParent_autogen(char_table lispObject) (lispObject, error) {
	return ec.stub("char-table-parent")
}

func (ec *execContext) setCharTableParent_autogen(char_table, parent lispObject) (lispObject, error) {
	return ec.stub("set-char-table-parent")
}

func (ec *execContext) charTableExtraSlot_autogen(char_table, n lispObject) (lispObject, error) {
	return ec.stub("char-table-extra-slot")
}

func (ec *execContext) setCharTableExtraSlot_autogen(char_table, n, value lispObject) (lispObject, error) {
	return ec.stub("set-char-table-extra-slot")
}

func (ec *execContext) charTableRange_autogen(char_table, range_ lispObject) (lispObject, error) {
	return ec.stub("char-table-range")
}

func (ec *execContext) setCharTableRange_autogen(char_table, range_, value lispObject) (lispObject, error) {
	return ec.stub("set-char-table-range")
}

func (ec *execContext) optimizeCharTable_autogen(char_table, test lispObject) (lispObject, error) {
	return ec.stub("optimize-char-table")
}

func (ec *execContext) mapCharTable_autogen(function, char_table lispObject) (lispObject, error) {
	return ec.stub("map-char-table")
}

func (ec *execContext) unicodePropertyTableInternal_autogen(prop lispObject) (lispObject, error) {
	return ec.stub("unicode-property-table-internal")
}

func (ec *execContext) getUnicodePropertyInternal_autogen(char_table, ch lispObject) (lispObject, error) {
	return ec.stub("get-unicode-property-internal")
}

func (ec *execContext) putUnicodePropertyInternal_autogen(char_table, ch, value lispObject) (lispObject, error) {
	return ec.stub("put-unicode-property-internal")
}

func (ec *execContext) xSelectFont_autogen(frame, exclude_proportional lispObject) (lispObject, error) {
	return ec.stub("x-select-font")
}

func (ec *execContext) xSelectFont_1_autogen(frame, ignored lispObject) (lispObject, error) {
	return ec.stub("x-select-font")
}

func (ec *execContext) xSelectFont_2_autogen(frame, ignored lispObject) (lispObject, error) {
	return ec.stub("x-select-font")
}

func (ec *execContext) xSelectFont_3_autogen(frame, exclude_proportional lispObject) (lispObject, error) {
	return ec.stub("x-select-font")
}

func (ec *execContext) identity_autogen(argument lispObject) (lispObject, error) {
	return ec.stub("identity")
}

func (ec *execContext) random_autogen(limit lispObject) (lispObject, error) {
	return ec.stub("random")
}

func (ec *execContext) length_autogen(sequence lispObject) (lispObject, error) {
	return ec.stub("length")
}

func (ec *execContext) safeLength_autogen(list lispObject) (lispObject, error) {
	return ec.stub("safe-length")
}

func (ec *execContext) lengthLess_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length<")
}

func (ec *execContext) lengthGreater_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length>")
}

func (ec *execContext) lengthEqual_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length=")
}

func (ec *execContext) properListP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("proper-list-p")
}

func (ec *execContext) stringBytes_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-bytes")
}

func (ec *execContext) stringDistance_autogen(string1, string2, bytecompare lispObject) (lispObject, error) {
	return ec.stub("string-distance")
}

func (ec *execContext) stringEqual_autogen(s1, s2 lispObject) (lispObject, error) {
	return ec.stub("string-equal")
}

func (ec *execContext) compareStrings_autogen(str1, start1, end1, str2, start2, end2, ignore_case lispObject) (lispObject, error) {
	return ec.stub("compare-strings")
}

func (ec *execContext) stringLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return ec.stub("string-lessp")
}

func (ec *execContext) stringVersionLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return ec.stub("string-version-lessp")
}

func (ec *execContext) stringCollateLessp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return ec.stub("string-collate-lessp")
}

func (ec *execContext) stringCollateEqualp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return ec.stub("string-collate-equalp")
}

func (ec *execContext) append_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("append")
}

func (ec *execContext) concat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("concat")
}

func (ec *execContext) vconcat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("vconcat")
}

func (ec *execContext) copySequence_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("copy-sequence")
}

func (ec *execContext) stringMakeMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-make-multibyte")
}

func (ec *execContext) stringMakeUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-make-unibyte")
}

func (ec *execContext) stringAsUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-as-unibyte")
}

func (ec *execContext) stringAsMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-as-multibyte")
}

func (ec *execContext) stringToMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-multibyte")
}

func (ec *execContext) stringToUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-unibyte")
}

func (ec *execContext) copyAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("copy-alist")
}

func (ec *execContext) substring_autogen(string, from, to lispObject) (lispObject, error) {
	return ec.stub("substring")
}

func (ec *execContext) substringNoProperties_autogen(string, from, to lispObject) (lispObject, error) {
	return ec.stub("substring-no-properties")
}

func (ec *execContext) take_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("take")
}

func (ec *execContext) ntake_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("ntake")
}

func (ec *execContext) nthcdr_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("nthcdr")
}

func (ec *execContext) nth_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("nth")
}

func (ec *execContext) elt_autogen(sequence, n lispObject) (lispObject, error) {
	return ec.stub("elt")
}

func (ec *execContext) member_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("member")
}

func (ec *execContext) memq_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("memq")
}

func (ec *execContext) memql_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("memql")
}

func (ec *execContext) assq_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("assq")
}

func (ec *execContext) assoc_autogen(key, alist, testfn lispObject) (lispObject, error) {
	return ec.stub("assoc")
}

func (ec *execContext) rassq_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("rassq")
}

func (ec *execContext) rassoc_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("rassoc")
}

func (ec *execContext) delq_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("delq")
}

func (ec *execContext) delete_autogen(elt, seq lispObject) (lispObject, error) {
	return ec.stub("delete")
}

func (ec *execContext) nreverse_autogen(seq lispObject) (lispObject, error) {
	return ec.stub("nreverse")
}

func (ec *execContext) reverse_autogen(seq lispObject) (lispObject, error) {
	return ec.stub("reverse")
}

func (ec *execContext) sort_autogen(seq, predicate lispObject) (lispObject, error) {
	return ec.stub("sort")
}

func (ec *execContext) plistGet_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-get")
}

func (ec *execContext) get_autogen(symbol, propname lispObject) (lispObject, error) {
	return ec.stub("get")
}

func (ec *execContext) plistPut_autogen(plist, prop, val, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-put")
}

func (ec *execContext) put_autogen(symbol, propname, value lispObject) (lispObject, error) {
	return ec.stub("put")
}

func (ec *execContext) plistMember_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-member")
}

func (ec *execContext) eql_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return ec.stub("eql")
}

func (ec *execContext) equal_autogen(o1, o2 lispObject) (lispObject, error) {
	return ec.stub("equal")
}

func (ec *execContext) equalIncludingProperties_autogen(o1, o2 lispObject) (lispObject, error) {
	return ec.stub("equal-including-properties")
}

func (ec *execContext) fillarray_autogen(array, item lispObject) (lispObject, error) {
	return ec.stub("fillarray")
}

func (ec *execContext) clearString_autogen(string lispObject) (lispObject, error) {
	return ec.stub("clear-string")
}

func (ec *execContext) nconc_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("nconc")
}

func (ec *execContext) mapconcat_autogen(function, sequence, separator lispObject) (lispObject, error) {
	return ec.stub("mapconcat")
}

func (ec *execContext) mapcar_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapcar")
}

func (ec *execContext) mapc_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapc")
}

func (ec *execContext) mapcan_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapcan")
}

func (ec *execContext) yesOrNoP_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("yes-or-no-p")
}

func (ec *execContext) loadAverage_autogen(use_floats lispObject) (lispObject, error) {
	return ec.stub("load-average")
}

func (ec *execContext) featurep_autogen(feature, subfeature lispObject) (lispObject, error) {
	return ec.stub("featurep")
}

func (ec *execContext) provide_autogen(feature, subfeatures lispObject) (lispObject, error) {
	return ec.stub("provide")
}

func (ec *execContext) require_autogen(feature, filename, noerror lispObject) (lispObject, error) {
	return ec.stub("require")
}

func (ec *execContext) widgetPut_autogen(widget, property, value lispObject) (lispObject, error) {
	return ec.stub("widget-put")
}

func (ec *execContext) widgetGet_autogen(widget, property lispObject) (lispObject, error) {
	return ec.stub("widget-get")
}

func (ec *execContext) widgetApply_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("widget-apply")
}

func (ec *execContext) localeInfo_autogen(item lispObject) (lispObject, error) {
	return ec.stub("locale-info")
}

func (ec *execContext) base64EncodeRegion_autogen(beg, end, no_line_break lispObject) (lispObject, error) {
	return ec.stub("base64-encode-region")
}

func (ec *execContext) base64urlEncodeRegion_autogen(beg, end, no_pad lispObject) (lispObject, error) {
	return ec.stub("base64url-encode-region")
}

func (ec *execContext) base64EncodeString_autogen(string, no_line_break lispObject) (lispObject, error) {
	return ec.stub("base64-encode-string")
}

func (ec *execContext) base64urlEncodeString_autogen(string, no_pad lispObject) (lispObject, error) {
	return ec.stub("base64url-encode-string")
}

func (ec *execContext) base64DecodeRegion_autogen(beg, end, base64url, ignore_invalid lispObject) (lispObject, error) {
	return ec.stub("base64-decode-region")
}

func (ec *execContext) base64DecodeString_autogen(string, base64url, ignore_invalid lispObject) (lispObject, error) {
	return ec.stub("base64-decode-string")
}

func (ec *execContext) sxhashEq_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-eq")
}

func (ec *execContext) sxhashEql_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-eql")
}

func (ec *execContext) sxhashEqual_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-equal")
}

func (ec *execContext) sxhashEqualIncludingProperties_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-equal-including-properties")
}

func (ec *execContext) makeHashTable_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-hash-table")
}

func (ec *execContext) copyHashTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-hash-table")
}

func (ec *execContext) hashTableCount_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-count")
}

func (ec *execContext) hashTableRehashSize_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-rehash-size")
}

func (ec *execContext) hashTableRehashThreshold_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-rehash-threshold")
}

func (ec *execContext) hashTableSize_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-size")
}

func (ec *execContext) hashTableTest_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-test")
}

func (ec *execContext) hashTableWeakness_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-weakness")
}

func (ec *execContext) hashTableP_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("hash-table-p")
}

func (ec *execContext) clrhash_autogen(table lispObject) (lispObject, error) {
	return ec.stub("clrhash")
}

func (ec *execContext) gethash_autogen(key, table, dflt lispObject) (lispObject, error) {
	return ec.stub("gethash")
}

func (ec *execContext) puthash_autogen(key, value, table lispObject) (lispObject, error) {
	return ec.stub("puthash")
}

func (ec *execContext) remhash_autogen(key, table lispObject) (lispObject, error) {
	return ec.stub("remhash")
}

func (ec *execContext) maphash_autogen(function, table lispObject) (lispObject, error) {
	return ec.stub("maphash")
}

func (ec *execContext) defineHashTableTest_autogen(name, test, hash lispObject) (lispObject, error) {
	return ec.stub("define-hash-table-test")
}

func (ec *execContext) secureHashAlgorithms_autogen() (lispObject, error) {
	return ec.stub("secure-hash-algorithms")
}

func (ec *execContext) md5_autogen(object, start, end, coding_system, noerror lispObject) (lispObject, error) {
	return ec.stub("md5")
}

func (ec *execContext) secureHash_autogen(algorithm, object, start, end, binary lispObject) (lispObject, error) {
	return ec.stub("secure-hash")
}

func (ec *execContext) bufferHash_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("buffer-hash")
}

func (ec *execContext) bufferLineStatistics_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("buffer-line-statistics")
}

func (ec *execContext) stringSearch_autogen(needle, haystack, start_pos lispObject) (lispObject, error) {
	return ec.stub("string-search")
}

func (ec *execContext) objectIntervals_autogen(object lispObject) (lispObject, error) {
	return ec.stub("object-intervals")
}

func (ec *execContext) lineNumberAtPos_autogen(position, absolute lispObject) (lispObject, error) {
	return ec.stub("line-number-at-pos")
}

func (ec *execContext) handleSaveSession_autogen(event lispObject) (lispObject, error) {
	return ec.stub("handle-save-session")
}

func (ec *execContext) setBufferRedisplay_autogen(symbol, newval, op, where lispObject) (lispObject, error) {
	return ec.stub("set-buffer-redisplay")
}

func (ec *execContext) linePixelHeight_autogen() (lispObject, error) {
	return ec.stub("line-pixel-height")
}

func (ec *execContext) getDisplayProperty_autogen(position, prop, object, properties lispObject) (lispObject, error) {
	return ec.stub("get-display-property")
}

func (ec *execContext) windowTextPixelSize_autogen(window, from, to, x_limit, y_limit, mode_lines, ignore_line_at_end lispObject) (lispObject, error) {
	return ec.stub("window-text-pixel-size")
}

func (ec *execContext) bufferTextPixelSize_autogen(buffer_or_name, window, x_limit, y_limit lispObject) (lispObject, error) {
	return ec.stub("buffer-text-pixel-size")
}

func (ec *execContext) displayLineIsContinuedP_autogen() (lispObject, error) {
	return ec.stub("display--line-is-continued-p")
}

func (ec *execContext) tabBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return ec.stub("tab-bar-height")
}

func (ec *execContext) toolBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return ec.stub("tool-bar-height")
}

func (ec *execContext) longLineOptimizationsP_autogen() (lispObject, error) {
	return ec.stub("long-line-optimizations-p")
}

func (ec *execContext) dumpGlyphMatrix_autogen(glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-glyph-matrix")
}

func (ec *execContext) dumpFrameGlyphMatrix_autogen() (lispObject, error) {
	return ec.stub("dump-frame-glyph-matrix")
}

func (ec *execContext) dumpGlyphRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-glyph-row")
}

func (ec *execContext) dumpTabBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-tab-bar-row")
}

func (ec *execContext) dumpToolBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-tool-bar-row")
}

func (ec *execContext) traceRedisplay_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("trace-redisplay")
}

func (ec *execContext) traceToStderr_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("trace-to-stderr")
}

func (ec *execContext) currentBidiParagraphDirection_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("current-bidi-paragraph-direction")
}

func (ec *execContext) bidiFindOverriddenDirectionality_autogen(from, to, object, base_dir lispObject) (lispObject, error) {
	return ec.stub("bidi-find-overridden-directionality")
}

func (ec *execContext) movePointVisually_autogen(direction lispObject) (lispObject, error) {
	return ec.stub("move-point-visually")
}

func (ec *execContext) bidiResolvedLevels_autogen(vpos lispObject) (lispObject, error) {
	return ec.stub("bidi-resolved-levels")
}

func (ec *execContext) formatModeLine_autogen(format, face, window, buffer lispObject) (lispObject, error) {
	return ec.stub("format-mode-line")
}

func (ec *execContext) invisibleP_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("invisible-p")
}

func (ec *execContext) lookupImageMap_autogen(map_, x, y lispObject) (lispObject, error) {
	return ec.stub("lookup-image-map")
}

func (ec *execContext) documentation_autogen(function, raw lispObject) (lispObject, error) {
	return ec.stub("documentation")
}

func (ec *execContext) documentationProperty_autogen(symbol, prop, raw lispObject) (lispObject, error) {
	return ec.stub("documentation-property")
}

func (ec *execContext) snarfDocumentation_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("Snarf-documentation")
}

func (ec *execContext) textQuotingStyle_autogen() (lispObject, error) {
	return ec.stub("text-quoting-style")
}

func (ec *execContext) byteCode_autogen(bytestr, vector, maxdepth lispObject) (lispObject, error) {
	return ec.stub("byte-code")
}

func (ec *execContext) internalStackStats_autogen() (lispObject, error) {
	return ec.stub("internal-stack-stats")
}

func (ec *execContext) queryFontset_autogen(pattern, regexpp lispObject) (lispObject, error) {
	return ec.stub("query-fontset")
}

func (ec *execContext) setFontsetFont_autogen(fontset, characters, font_spec, frame, add lispObject) (lispObject, error) {
	return ec.stub("set-fontset-font")
}

func (ec *execContext) newFontset_autogen(name, fontlist lispObject) (lispObject, error) {
	return ec.stub("new-fontset")
}

func (ec *execContext) fontsetInfo_autogen(fontset, frame lispObject) (lispObject, error) {
	return ec.stub("fontset-info")
}

func (ec *execContext) fontsetFont_autogen(name, ch, all lispObject) (lispObject, error) {
	return ec.stub("fontset-font")
}

func (ec *execContext) fontsetList_autogen() (lispObject, error) {
	return ec.stub("fontset-list")
}

func (ec *execContext) fontsetListAll_autogen() (lispObject, error) {
	return ec.stub("fontset-list-all")
}

func (ec *execContext) bufferLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("buffer-live-p")
}

func (ec *execContext) bufferList_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("buffer-list")
}

func (ec *execContext) getBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("get-buffer")
}

func (ec *execContext) getFileBuffer_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("get-file-buffer")
}

func (ec *execContext) getBufferCreate_autogen(buffer_or_name, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return ec.stub("get-buffer-create")
}

func (ec *execContext) makeIndirectBuffer_autogen(base_buffer, name, clone, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return ec.stub("make-indirect-buffer")
}

func (ec *execContext) generateNewBufferName_autogen(name, ignore lispObject) (lispObject, error) {
	return ec.stub("generate-new-buffer-name")
}

func (ec *execContext) bufferName_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-name")
}

func (ec *execContext) bufferFileName_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-file-name")
}

func (ec *execContext) bufferBaseBuffer_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-base-buffer")
}

func (ec *execContext) bufferLocalValue_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-local-value")
}

func (ec *execContext) bufferLocalVariables_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-local-variables")
}

func (ec *execContext) bufferModifiedP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-modified-p")
}

func (ec *execContext) forceModeLineUpdate_autogen(all lispObject) (lispObject, error) {
	return ec.stub("force-mode-line-update")
}

func (ec *execContext) setBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("set-buffer-modified-p")
}

func (ec *execContext) restoreBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("restore-buffer-modified-p")
}

func (ec *execContext) bufferModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-modified-tick")
}

func (ec *execContext) internalSetBufferModifiedTick_autogen(tick, buffer lispObject) (lispObject, error) {
	return ec.stub("internal--set-buffer-modified-tick")
}

func (ec *execContext) bufferCharsModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-chars-modified-tick")
}

func (ec *execContext) renameBuffer_autogen(newname, unique lispObject) (lispObject, error) {
	return ec.stub("rename-buffer")
}

func (ec *execContext) otherBuffer_autogen(buffer, visible_ok, frame lispObject) (lispObject, error) {
	return ec.stub("other-buffer")
}

func (ec *execContext) bufferEnableUndo_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-enable-undo")
}

func (ec *execContext) killBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("kill-buffer")
}

func (ec *execContext) buryBufferInternal_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("bury-buffer-internal")
}

func (ec *execContext) setBufferMajorMode_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("set-buffer-major-mode")
}

func (ec *execContext) currentBuffer_autogen() (lispObject, error) {
	return ec.stub("current-buffer")
}

func (ec *execContext) setBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("set-buffer")
}

func (ec *execContext) barfIfBufferReadOnly_autogen(position lispObject) (lispObject, error) {
	return ec.stub("barf-if-buffer-read-only")
}

func (ec *execContext) eraseBuffer_autogen() (lispObject, error) {
	return ec.stub("erase-buffer")
}

func (ec *execContext) bufferSwapText_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-swap-text")
}

func (ec *execContext) setBufferMultibyte_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("set-buffer-multibyte")
}

func (ec *execContext) killAllLocalVariables_autogen(kill_permanent lispObject) (lispObject, error) {
	return ec.stub("kill-all-local-variables")
}

func (ec *execContext) overlayp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("overlayp")
}

func (ec *execContext) makeOverlay_autogen(beg, end, buffer, front_advance, rear_advance lispObject) (lispObject, error) {
	return ec.stub("make-overlay")
}

func (ec *execContext) moveOverlay_autogen(overlay, beg, end, buffer lispObject) (lispObject, error) {
	return ec.stub("move-overlay")
}

func (ec *execContext) deleteOverlay_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("delete-overlay")
}

func (ec *execContext) deleteAllOverlays_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("delete-all-overlays")
}

func (ec *execContext) overlayStart_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-start")
}

func (ec *execContext) overlayEnd_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-end")
}

func (ec *execContext) overlayBuffer_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-buffer")
}

func (ec *execContext) overlayProperties_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-properties")
}

func (ec *execContext) overlaysAt_autogen(pos, sorted lispObject) (lispObject, error) {
	return ec.stub("overlays-at")
}

func (ec *execContext) overlaysIn_autogen(beg, end lispObject) (lispObject, error) {
	return ec.stub("overlays-in")
}

func (ec *execContext) nextOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("next-overlay-change")
}

func (ec *execContext) previousOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("previous-overlay-change")
}

func (ec *execContext) overlayLists_autogen() (lispObject, error) {
	return ec.stub("overlay-lists")
}

func (ec *execContext) overlayRecenter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("overlay-recenter")
}

func (ec *execContext) overlayGet_autogen(overlay, prop lispObject) (lispObject, error) {
	return ec.stub("overlay-get")
}

func (ec *execContext) overlayPut_autogen(overlay, prop, value lispObject) (lispObject, error) {
	return ec.stub("overlay-put")
}

func (ec *execContext) overlayTree_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("overlay-tree")
}

func (ec *execContext) forwardChar_autogen(n lispObject) (lispObject, error) {
	return ec.stub("forward-char")
}

func (ec *execContext) backwardChar_autogen(n lispObject) (lispObject, error) {
	return ec.stub("backward-char")
}

func (ec *execContext) forwardLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("forward-line")
}

func (ec *execContext) beginningOfLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("beginning-of-line")
}

func (ec *execContext) endOfLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("end-of-line")
}

func (ec *execContext) deleteChar_autogen(n, killflag lispObject) (lispObject, error) {
	return ec.stub("delete-char")
}

func (ec *execContext) selfInsertCommand_autogen(n, c lispObject) (lispObject, error) {
	return ec.stub("self-insert-command")
}

func (ec *execContext) readChar_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-char")
}

func (ec *execContext) readEvent_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-event")
}

func (ec *execContext) readCharExclusive_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-char-exclusive")
}

func (ec *execContext) getFileChar_autogen() (lispObject, error) {
	return ec.stub("get-file-char")
}

func (ec *execContext) getLoadSuffixes_autogen() (lispObject, error) {
	return ec.stub("get-load-suffixes")
}

func (ec *execContext) load_autogen(file, noerror, nomessage, nosuffix, must_suffix lispObject) (lispObject, error) {
	return ec.stub("load")
}

func (ec *execContext) locateFileInternal_autogen(filename, path, suffixes, predicate lispObject) (lispObject, error) {
	return ec.stub("locate-file-internal")
}

func (ec *execContext) evalBuffer_autogen(buffer, printflag, filename, unibyte, do_allow_print lispObject) (lispObject, error) {
	return ec.stub("eval-buffer")
}

func (ec *execContext) evalRegion_autogen(start, end, printflag, read_function lispObject) (lispObject, error) {
	return ec.stub("eval-region")
}

func (ec *execContext) read_autogen(stream lispObject) (lispObject, error) {
	return ec.stub("read")
}

func (ec *execContext) readPositioningSymbols_autogen(stream lispObject) (lispObject, error) {
	return ec.stub("read-positioning-symbols")
}

func (ec *execContext) readFromString_autogen(string, start, end lispObject) (lispObject, error) {
	return ec.stub("read-from-string")
}

func (ec *execContext) lreadSubstituteObjectInSubtree_autogen(object, placeholder, completed lispObject) (lispObject, error) {
	return ec.stub("lread--substitute-object-in-subtree")
}

func (ec *execContext) intern_autogen(string, obarray lispObject) (lispObject, error) {
	return ec.stub("intern")
}

func (ec *execContext) internSoft_autogen(name, obarray lispObject) (lispObject, error) {
	return ec.stub("intern-soft")
}

func (ec *execContext) unintern_autogen(name, obarray lispObject) (lispObject, error) {
	return ec.stub("unintern")
}

func (ec *execContext) mapatoms_autogen(function, obarray lispObject) (lispObject, error) {
	return ec.stub("mapatoms")
}

func (ec *execContext) pgtkUseImContext_autogen(use_p, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-use-im-context")
}

func (ec *execContext) getInternalRunTime_autogen() (lispObject, error) {
	return ec.stub("get-internal-run-time")
}

func (ec *execContext) callProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("call-process")
}

func (ec *execContext) callProcessRegion_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("call-process-region")
}

func (ec *execContext) getenvInternal_autogen(variable, env lispObject) (lispObject, error) {
	return ec.stub("getenv-internal")
}

func (ec *execContext) findFileNameHandler_autogen(filename, operation lispObject) (lispObject, error) {
	return ec.stub("find-file-name-handler")
}

func (ec *execContext) fileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-directory")
}

func (ec *execContext) fileNameNondirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-nondirectory")
}

func (ec *execContext) unhandledFileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("unhandled-file-name-directory")
}

func (ec *execContext) fileNameAsDirectory_autogen(file lispObject) (lispObject, error) {
	return ec.stub("file-name-as-directory")
}

func (ec *execContext) directoryNameP_autogen(name lispObject) (lispObject, error) {
	return ec.stub("directory-name-p")
}

func (ec *execContext) directoryFileName_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("directory-file-name")
}

func (ec *execContext) makeTempFileInternal_autogen(prefix, dir_flag, suffix, text lispObject) (lispObject, error) {
	return ec.stub("make-temp-file-internal")
}

func (ec *execContext) makeTempName_autogen(prefix lispObject) (lispObject, error) {
	return ec.stub("make-temp-name")
}

func (ec *execContext) fileNameConcat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("file-name-concat")
}

func (ec *execContext) expandFileName_autogen(name, default_directory lispObject) (lispObject, error) {
	return ec.stub("expand-file-name")
}

func (ec *execContext) substituteInFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("substitute-in-file-name")
}

func (ec *execContext) copyFile_autogen(file, newname, ok_if_already_exists, keep_time, preserve_uid_gid, preserve_permissions lispObject) (lispObject, error) {
	return ec.stub("copy-file")
}

func (ec *execContext) makeDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("make-directory-internal")
}

func (ec *execContext) deleteDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("delete-directory-internal")
}

func (ec *execContext) deleteFile_autogen(filename, trash lispObject) (lispObject, error) {
	return ec.stub("delete-file")
}

func (ec *execContext) fileNameCaseInsensitiveP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-case-insensitive-p")
}

func (ec *execContext) renameFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("rename-file")
}

func (ec *execContext) addNameToFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("add-name-to-file")
}

func (ec *execContext) makeSymbolicLink_autogen(target, linkname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("make-symbolic-link")
}

func (ec *execContext) fileNameAbsoluteP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-absolute-p")
}

func (ec *execContext) fileExistsP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-exists-p")
}

func (ec *execContext) fileExecutableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-executable-p")
}

func (ec *execContext) fileReadableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-readable-p")
}

func (ec *execContext) fileWritableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-writable-p")
}

func (ec *execContext) accessFile_autogen(filename, string lispObject) (lispObject, error) {
	return ec.stub("access-file")
}

func (ec *execContext) fileSymlinkP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-symlink-p")
}

func (ec *execContext) fileDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-directory-p")
}

func (ec *execContext) fileAccessibleDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-accessible-directory-p")
}

func (ec *execContext) fileRegularP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-regular-p")
}

func (ec *execContext) fileSelinuxContext_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-selinux-context")
}

func (ec *execContext) setFileSelinuxContext_autogen(filename, context lispObject) (lispObject, error) {
	return ec.stub("set-file-selinux-context")
}

func (ec *execContext) fileAcl_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-acl")
}

func (ec *execContext) setFileAcl_autogen(filename, acl_string lispObject) (lispObject, error) {
	return ec.stub("set-file-acl")
}

func (ec *execContext) fileModes_autogen(filename, flag lispObject) (lispObject, error) {
	return ec.stub("file-modes")
}

func (ec *execContext) setFileModes_autogen(filename, mode, flag lispObject) (lispObject, error) {
	return ec.stub("set-file-modes")
}

func (ec *execContext) setDefaultFileModes_autogen(mode lispObject) (lispObject, error) {
	return ec.stub("set-default-file-modes")
}

func (ec *execContext) defaultFileModes_autogen() (lispObject, error) {
	return ec.stub("default-file-modes")
}

func (ec *execContext) setFileTimes_autogen(filename, timestamp, flag lispObject) (lispObject, error) {
	return ec.stub("set-file-times")
}

func (ec *execContext) unixSync_autogen() (lispObject, error) {
	return ec.stub("unix-sync")
}

func (ec *execContext) fileNewerThanFileP_autogen(file1, file2 lispObject) (lispObject, error) {
	return ec.stub("file-newer-than-file-p")
}

func (ec *execContext) insertFileContents_autogen(filename, visit, beg, end, replace lispObject) (lispObject, error) {
	return ec.stub("insert-file-contents")
}

func (ec *execContext) writeRegion_autogen(start, end, filename, append, visit, lockname, mustbenew lispObject) (lispObject, error) {
	return ec.stub("write-region")
}

func (ec *execContext) carLessThanCar_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("car-less-than-car")
}

func (ec *execContext) verifyVisitedFileModtime_autogen(buf lispObject) (lispObject, error) {
	return ec.stub("verify-visited-file-modtime")
}

func (ec *execContext) visitedFileModtime_autogen() (lispObject, error) {
	return ec.stub("visited-file-modtime")
}

func (ec *execContext) setVisitedFileModtime_autogen(time_flag lispObject) (lispObject, error) {
	return ec.stub("set-visited-file-modtime")
}

func (ec *execContext) doAutoSave_autogen(no_message, current_only lispObject) (lispObject, error) {
	return ec.stub("do-auto-save")
}

func (ec *execContext) setBufferAutoSaved_autogen() (lispObject, error) {
	return ec.stub("set-buffer-auto-saved")
}

func (ec *execContext) clearBufferAutoSaveFailure_autogen() (lispObject, error) {
	return ec.stub("clear-buffer-auto-save-failure")
}

func (ec *execContext) recentAutoSaveP_autogen() (lispObject, error) {
	return ec.stub("recent-auto-save-p")
}

func (ec *execContext) nextReadFileUsesDialogP_autogen() (lispObject, error) {
	return ec.stub("next-read-file-uses-dialog-p")
}

func (ec *execContext) setBinaryMode_autogen(stream, mode lispObject) (lispObject, error) {
	return ec.stub("set-binary-mode")
}

func (ec *execContext) haikuSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("haiku-set-mouse-absolute-pixel-position")
}

func (ec *execContext) haikuMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("haiku-mouse-absolute-pixel-position")
}

func (ec *execContext) xwDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p")
}

func (ec *execContext) xwDisplayColorP_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p")
}

func (ec *execContext) xwDisplayColorP_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p")
}

func (ec *execContext) xwDisplayColorP_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p")
}

func (ec *execContext) xwColorDefinedP_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p")
}

func (ec *execContext) xwColorDefinedP_1_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p")
}

func (ec *execContext) xwColorDefinedP_2_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p")
}

func (ec *execContext) xwColorDefinedP_3_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p")
}

func (ec *execContext) xwColorValues_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values")
}

func (ec *execContext) xwColorValues_1_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values")
}

func (ec *execContext) xwColorValues_2_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values")
}

func (ec *execContext) xwColorValues_3_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values")
}

func (ec *execContext) xDisplayGrayscaleP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p")
}

func (ec *execContext) xDisplayGrayscaleP_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p")
}

func (ec *execContext) xDisplayGrayscaleP_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p")
}

func (ec *execContext) xDisplayGrayscaleP_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p")
}

func (ec *execContext) xOpenConnection_autogen(display, resource_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection")
}

func (ec *execContext) xOpenConnection_1_autogen(display, resource_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection")
}

func (ec *execContext) xOpenConnection_2_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection")
}

func (ec *execContext) xOpenConnection_3_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection")
}

func (ec *execContext) xDisplayPixelWidth_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width")
}

func (ec *execContext) xDisplayPixelWidth_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width")
}

func (ec *execContext) xDisplayPixelWidth_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width")
}

func (ec *execContext) xDisplayPixelWidth_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width")
}

func (ec *execContext) xDisplayPixelHeight_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height")
}

func (ec *execContext) xDisplayPixelHeight_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height")
}

func (ec *execContext) xDisplayPixelHeight_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height")
}

func (ec *execContext) xDisplayPixelHeight_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height")
}

func (ec *execContext) xDisplayMmHeight_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height")
}

func (ec *execContext) xDisplayMmHeight_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height")
}

func (ec *execContext) xDisplayMmHeight_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height")
}

func (ec *execContext) xDisplayMmHeight_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height")
}

func (ec *execContext) xDisplayMmWidth_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width")
}

func (ec *execContext) xDisplayMmWidth_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width")
}

func (ec *execContext) xDisplayMmWidth_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width")
}

func (ec *execContext) xDisplayMmWidth_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width")
}

func (ec *execContext) xCreateFrame_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame")
}

func (ec *execContext) xCreateFrame_1_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame")
}

func (ec *execContext) xCreateFrame_2_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame")
}

func (ec *execContext) xCreateFrame_3_autogen(parameters lispObject) (lispObject, error) {
	return ec.stub("x-create-frame")
}

func (ec *execContext) xDisplayVisualClass_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class")
}

func (ec *execContext) xDisplayVisualClass_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class")
}

func (ec *execContext) xDisplayVisualClass_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class")
}

func (ec *execContext) xDisplayVisualClass_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class")
}

func (ec *execContext) xShowTip_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip")
}

func (ec *execContext) xShowTip_1_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip")
}

func (ec *execContext) xShowTip_2_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip")
}

func (ec *execContext) xShowTip_3_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip")
}

func (ec *execContext) xHideTip_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip")
}

func (ec *execContext) xHideTip_1_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip")
}

func (ec *execContext) xHideTip_2_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip")
}

func (ec *execContext) xHideTip_3_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip")
}

func (ec *execContext) xCloseConnection_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection")
}

func (ec *execContext) xCloseConnection_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection")
}

func (ec *execContext) xCloseConnection_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection")
}

func (ec *execContext) xCloseConnection_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-close-connection")
}

func (ec *execContext) xDisplayList_autogen() (lispObject, error) {
	return ec.stub("x-display-list")
}

func (ec *execContext) xDisplayList_1_autogen() (lispObject, error) {
	return ec.stub("x-display-list")
}

func (ec *execContext) xDisplayList_2_autogen() (lispObject, error) {
	return ec.stub("x-display-list")
}

func (ec *execContext) xDisplayList_3_autogen() (lispObject, error) {
	return ec.stub("x-display-list")
}

func (ec *execContext) xServerVendor_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor")
}

func (ec *execContext) xServerVendor_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor")
}

func (ec *execContext) xServerVendor_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor")
}

func (ec *execContext) xServerVersion_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version")
}

func (ec *execContext) xServerVersion_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version")
}

func (ec *execContext) xServerVersion_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version")
}

func (ec *execContext) xDisplayScreens_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens")
}

func (ec *execContext) xDisplayScreens_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens")
}

func (ec *execContext) xDisplayScreens_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens")
}

func (ec *execContext) xDisplayScreens_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-screens")
}

func (ec *execContext) haikuGetVersionString_autogen() (lispObject, error) {
	return ec.stub("haiku-get-version-string")
}

func (ec *execContext) xDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells")
}

func (ec *execContext) xDisplayColorCells_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells")
}

func (ec *execContext) xDisplayColorCells_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells")
}

func (ec *execContext) xDisplayColorCells_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells")
}

func (ec *execContext) xDisplayPlanes_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes")
}

func (ec *execContext) xDisplayPlanes_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes")
}

func (ec *execContext) xDisplayPlanes_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes")
}

func (ec *execContext) xDisplayPlanes_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-planes")
}

func (ec *execContext) xDoubleBufferedP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-double-buffered-p")
}

func (ec *execContext) xDoubleBufferedP_1_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-double-buffered-p")
}

func (ec *execContext) xDisplayBackingStore_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store")
}

func (ec *execContext) xDisplayBackingStore_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store")
}

func (ec *execContext) xDisplayBackingStore_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store")
}

func (ec *execContext) xDisplayBackingStore_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store")
}

func (ec *execContext) haikuFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-geometry")
}

func (ec *execContext) haikuFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-edges")
}

func (ec *execContext) haikuReadFileName_autogen(prompt, frame, dir, mustmatch, dir_only_p, save_text lispObject) (lispObject, error) {
	return ec.stub("haiku-read-file-name")
}

func (ec *execContext) haikuPutResource_autogen(resource, string lispObject) (lispObject, error) {
	return ec.stub("haiku-put-resource")
}

func (ec *execContext) haikuFrameListZOrder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-list-z-order")
}

func (ec *execContext) xDisplaySaveUnder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under")
}

func (ec *execContext) xDisplaySaveUnder_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under")
}

func (ec *execContext) xDisplaySaveUnder_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under")
}

func (ec *execContext) xDisplaySaveUnder_3_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under")
}

func (ec *execContext) haikuFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-restack")
}

func (ec *execContext) haikuSaveSessionReply_autogen(quit_reply lispObject) (lispObject, error) {
	return ec.stub("haiku-save-session-reply")
}

func (ec *execContext) haikuDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("haiku-display-monitor-attributes-list")
}

func (ec *execContext) xOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return ec.stub("x-own-selection-internal")
}

func (ec *execContext) xGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return ec.stub("x-get-selection-internal")
}

func (ec *execContext) xDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("x-disown-selection-internal")
}

func (ec *execContext) xSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("x-selection-owner-p")
}

func (ec *execContext) xSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("x-selection-exists-p")
}

func (ec *execContext) xGetLocalSelection_autogen(value, target lispObject) (lispObject, error) {
	return ec.stub("x-get-local-selection")
}

func (ec *execContext) xGetAtomName_autogen(value, frame lispObject) (lispObject, error) {
	return ec.stub("x-get-atom-name")
}

func (ec *execContext) xRegisterDndAtom_autogen(atom, frame lispObject) (lispObject, error) {
	return ec.stub("x-register-dnd-atom")
}

func (ec *execContext) xSendClientMessage_autogen(display, dest, from, message_type, format, values lispObject) (lispObject, error) {
	return ec.stub("x-send-client-message")
}

func (ec *execContext) w32notifyAddWatch_autogen(file, filter, callback lispObject) (lispObject, error) {
	return ec.stub("w32notify-add-watch")
}

func (ec *execContext) w32notifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("w32notify-rm-watch")
}

func (ec *execContext) w32notifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("w32notify-valid-p")
}

func (ec *execContext) w32SetClipboardData_autogen(string, ignored lispObject) (lispObject, error) {
	return ec.stub("w32-set-clipboard-data")
}

func (ec *execContext) w32GetClipboardData_autogen(ignored lispObject) (lispObject, error) {
	return ec.stub("w32-get-clipboard-data")
}

func (ec *execContext) w32SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w32-selection-exists-p")
}

func (ec *execContext) w32SelectionTargets_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w32-selection-targets")
}

func (ec *execContext) gnutlsAsynchronousParameters_autogen(proc, params lispObject) (lispObject, error) {
	return ec.stub("gnutls-asynchronous-parameters")
}

func (ec *execContext) gnutlsGetInitstage_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-get-initstage")
}

func (ec *execContext) gnutlsErrorp_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-errorp")
}

func (ec *execContext) gnutlsErrorFatalp_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-error-fatalp")
}

func (ec *execContext) gnutlsErrorString_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-error-string")
}

func (ec *execContext) gnutlsDeinit_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-deinit")
}

func (ec *execContext) gnutlsPeerStatusWarningDescribe_autogen(status_symbol lispObject) (lispObject, error) {
	return ec.stub("gnutls-peer-status-warning-describe")
}

func (ec *execContext) gnutlsPeerStatus_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-peer-status")
}

func (ec *execContext) gnutlsFormatCertificate_autogen(cert lispObject) (lispObject, error) {
	return ec.stub("gnutls-format-certificate")
}

func (ec *execContext) gnutlsBoot_autogen(proc, type_, proplist lispObject) (lispObject, error) {
	return ec.stub("gnutls-boot")
}

func (ec *execContext) gnutlsBye_autogen(proc, cont lispObject) (lispObject, error) {
	return ec.stub("gnutls-bye")
}

func (ec *execContext) gnutlsCiphers_autogen() (lispObject, error) {
	return ec.stub("gnutls-ciphers")
}

func (ec *execContext) gnutlsSymmetricEncrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return ec.stub("gnutls-symmetric-encrypt")
}

func (ec *execContext) gnutlsSymmetricDecrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return ec.stub("gnutls-symmetric-decrypt")
}

func (ec *execContext) gnutlsMacs_autogen() (lispObject, error) {
	return ec.stub("gnutls-macs")
}

func (ec *execContext) gnutlsDigests_autogen() (lispObject, error) {
	return ec.stub("gnutls-digests")
}

func (ec *execContext) gnutlsHashMac_autogen(hash_method, key, input lispObject) (lispObject, error) {
	return ec.stub("gnutls-hash-mac")
}

func (ec *execContext) gnutlsHashDigest_autogen(digest_method, input lispObject) (lispObject, error) {
	return ec.stub("gnutls-hash-digest")
}

func (ec *execContext) gnutlsAvailableP_autogen() (lispObject, error) {
	return ec.stub("gnutls-available-p")
}

func (ec *execContext) lcmsCieDe2000_autogen(color1, color2, kL, kC, kH lispObject) (lispObject, error) {
	return ec.stub("lcms-cie-de2000")
}

func (ec *execContext) lcmsXyzToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-xyz->jch")
}

func (ec *execContext) lcmsJchToXyz_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jch->xyz")
}

func (ec *execContext) lcmsJchToJab_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jch->jab")
}

func (ec *execContext) lcmsJabToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jab->jch")
}

func (ec *execContext) lcmsCam02Ucs_autogen(color1, color2, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-cam02-ucs")
}

func (ec *execContext) lcmsTempToWhitePoint_autogen(temperature lispObject) (lispObject, error) {
	return ec.stub("lcms-temp->white-point")
}

func (ec *execContext) lcms2AvailableP_autogen() (lispObject, error) {
	return ec.stub("lcms2-available-p")
}

func (ec *execContext) codingSystemP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("coding-system-p")
}

func (ec *execContext) readNonNilCodingSystem_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("read-non-nil-coding-system")
}

func (ec *execContext) readCodingSystem_autogen(prompt, default_coding_system lispObject) (lispObject, error) {
	return ec.stub("read-coding-system")
}

func (ec *execContext) checkCodingSystem_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("check-coding-system")
}

func (ec *execContext) detectCodingRegion_autogen(start, end, highest lispObject) (lispObject, error) {
	return ec.stub("detect-coding-region")
}

func (ec *execContext) detectCodingString_autogen(string, highest lispObject) (lispObject, error) {
	return ec.stub("detect-coding-string")
}

func (ec *execContext) findCodingSystemsRegionInternal_autogen(start, end, exclude lispObject) (lispObject, error) {
	return ec.stub("find-coding-systems-region-internal")
}

func (ec *execContext) unencodableCharPosition_autogen(start, end, coding_system, count, string lispObject) (lispObject, error) {
	return ec.stub("unencodable-char-position")
}

func (ec *execContext) checkCodingSystemsRegion_autogen(start, end, coding_system_list lispObject) (lispObject, error) {
	return ec.stub("check-coding-systems-region")
}

func (ec *execContext) decodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return ec.stub("decode-coding-region")
}

func (ec *execContext) encodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return ec.stub("encode-coding-region")
}

func (ec *execContext) internalEncodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, encode_method, count lispObject) (lispObject, error) {
	return ec.stub("internal-encode-string-utf-8")
}

func (ec *execContext) internalDecodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, decode_method, count lispObject) (lispObject, error) {
	return ec.stub("internal-decode-string-utf-8")
}

func (ec *execContext) decodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return ec.stub("decode-coding-string")
}

func (ec *execContext) encodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return ec.stub("encode-coding-string")
}

func (ec *execContext) decodeSjisChar_autogen(code lispObject) (lispObject, error) {
	return ec.stub("decode-sjis-char")
}

func (ec *execContext) encodeSjisChar_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("encode-sjis-char")
}

func (ec *execContext) decodeBig5Char_autogen(code lispObject) (lispObject, error) {
	return ec.stub("decode-big5-char")
}

func (ec *execContext) encodeBig5Char_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("encode-big5-char")
}

func (ec *execContext) setTerminalCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return ec.stub("set-terminal-coding-system-internal")
}

func (ec *execContext) setSafeTerminalCodingSystemInternal_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("set-safe-terminal-coding-system-internal")
}

func (ec *execContext) terminalCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-coding-system")
}

func (ec *execContext) setKeyboardCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return ec.stub("set-keyboard-coding-system-internal")
}

func (ec *execContext) keyboardCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("keyboard-coding-system")
}

func (ec *execContext) findOperationCodingSystem_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("find-operation-coding-system")
}

func (ec *execContext) setCodingSystemPriority_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("set-coding-system-priority")
}

func (ec *execContext) codingSystemPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return ec.stub("coding-system-priority-list")
}

func (ec *execContext) defineCodingSystemInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("define-coding-system-internal")
}

func (ec *execContext) codingSystemPut_autogen(coding_system, prop, val lispObject) (lispObject, error) {
	return ec.stub("coding-system-put")
}

func (ec *execContext) defineCodingSystemAlias_autogen(alias, coding_system lispObject) (lispObject, error) {
	return ec.stub("define-coding-system-alias")
}

func (ec *execContext) codingSystemBase_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-base")
}

func (ec *execContext) codingSystemPlist_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-plist")
}

func (ec *execContext) codingSystemAliases_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-aliases")
}

func (ec *execContext) codingSystemEolType_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-eol-type")
}

func (ec *execContext) dumpColors_autogen() (lispObject, error) {
	return ec.stub("dump-colors")
}

func (ec *execContext) clearFaceCache_autogen(thoroughly lispObject) (lispObject, error) {
	return ec.stub("clear-face-cache")
}

func (ec *execContext) bitmapSpecP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bitmap-spec-p")
}

func (ec *execContext) colorValuesFromColorSpec_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("color-values-from-color-spec")
}

func (ec *execContext) colorGrayP_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("color-gray-p")
}

func (ec *execContext) colorSupportedP_autogen(color, frame, background_p lispObject) (lispObject, error) {
	return ec.stub("color-supported-p")
}

func (ec *execContext) xFamilyFonts_autogen(family, frame lispObject) (lispObject, error) {
	return ec.stub("x-family-fonts")
}

func (ec *execContext) xListFonts_autogen(pattern, face, frame, maximum, width lispObject) (lispObject, error) {
	return ec.stub("x-list-fonts")
}

func (ec *execContext) internalMakeLispFace_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-make-lisp-face")
}

func (ec *execContext) internalLispFaceP_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-p")
}

func (ec *execContext) internalCopyLispFace_autogen(from, to, frame, new_frame lispObject) (lispObject, error) {
	return ec.stub("internal-copy-lisp-face")
}

func (ec *execContext) internalSetLispFaceAttribute_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return ec.stub("internal-set-lisp-face-attribute")
}

func (ec *execContext) internalFaceXGetResource_autogen(resource, class, frame lispObject) (lispObject, error) {
	return ec.stub("internal-face-x-get-resource")
}

func (ec *execContext) internalSetLispFaceAttributeFromResource_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return ec.stub("internal-set-lisp-face-attribute-from-resource")
}

func (ec *execContext) faceAttributeRelativeP_autogen(attribute, value lispObject) (lispObject, error) {
	return ec.stub("face-attribute-relative-p")
}

func (ec *execContext) mergeFaceAttribute_autogen(attribute, value1, value2 lispObject) (lispObject, error) {
	return ec.stub("merge-face-attribute")
}

func (ec *execContext) internalGetLispFaceAttribute_autogen(symbol, keyword, frame lispObject) (lispObject, error) {
	return ec.stub("internal-get-lisp-face-attribute")
}

func (ec *execContext) internalLispFaceAttributeValues_autogen(attr lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-attribute-values")
}

func (ec *execContext) internalMergeInGlobalFace_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-merge-in-global-face")
}

func (ec *execContext) faceFont_autogen(face, frame, character lispObject) (lispObject, error) {
	return ec.stub("face-font")
}

func (ec *execContext) internalLispFaceEqualP_autogen(face1, face2, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-equal-p")
}

func (ec *execContext) internalLispFaceEmptyP_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-empty-p")
}

func (ec *execContext) frameFaceHashTable_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame--face-hash-table")
}

func (ec *execContext) colorDistance_autogen(color1, color2, frame, metric lispObject) (lispObject, error) {
	return ec.stub("color-distance")
}

func (ec *execContext) faceAttributesAsVector_autogen(plist lispObject) (lispObject, error) {
	return ec.stub("face-attributes-as-vector")
}

func (ec *execContext) displaySupportsFaceAttributesP_autogen(attributes, display lispObject) (lispObject, error) {
	return ec.stub("display-supports-face-attributes-p")
}

func (ec *execContext) internalSetFontSelectionOrder_autogen(order lispObject) (lispObject, error) {
	return ec.stub("internal-set-font-selection-order")
}

func (ec *execContext) internalSetAlternativeFontFamilyAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("internal-set-alternative-font-family-alist")
}

func (ec *execContext) internalSetAlternativeFontRegistryAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("internal-set-alternative-font-registry-alist")
}

func (ec *execContext) ttySuppressBoldInverseDefaultColors_autogen(suppress lispObject) (lispObject, error) {
	return ec.stub("tty-suppress-bold-inverse-default-colors")
}

func (ec *execContext) xLoadColorFile_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("x-load-color-file")
}

func (ec *execContext) dumpFace_autogen(n lispObject) (lispObject, error) {
	return ec.stub("dump-face")
}

func (ec *execContext) showFaceResources_autogen() (lispObject, error) {
	return ec.stub("show-face-resources")
}

func (ec *execContext) makeKeymap_autogen(string lispObject) (lispObject, error) {
	return ec.stub("make-keymap")
}

func (ec *execContext) makeSparseKeymap_autogen(string lispObject) (lispObject, error) {
	return ec.stub("make-sparse-keymap")
}

func (ec *execContext) keymapp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("keymapp")
}

func (ec *execContext) keymapPrompt_autogen(map_ lispObject) (lispObject, error) {
	return ec.stub("keymap-prompt")
}

func (ec *execContext) keymapParent_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("keymap-parent")
}

func (ec *execContext) setKeymapParent_autogen(keymap, parent lispObject) (lispObject, error) {
	return ec.stub("set-keymap-parent")
}

func (ec *execContext) mapKeymapInternal_autogen(function, keymap lispObject) (lispObject, error) {
	return ec.stub("map-keymap-internal")
}

func (ec *execContext) mapKeymap_autogen(function, keymap, sort_first lispObject) (lispObject, error) {
	return ec.stub("map-keymap")
}

func (ec *execContext) keymapGetKeyelt_autogen(object, autoload lispObject) (lispObject, error) {
	return ec.stub("keymap--get-keyelt")
}

func (ec *execContext) copyKeymap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("copy-keymap")
}

func (ec *execContext) defineKey_autogen(keymap, key, def, remove lispObject) (lispObject, error) {
	return ec.stub("define-key")
}

func (ec *execContext) commandRemapping_autogen(command, position, keymaps lispObject) (lispObject, error) {
	return ec.stub("command-remapping")
}

func (ec *execContext) lookupKey_autogen(keymap, key, accept_default lispObject) (lispObject, error) {
	return ec.stub("lookup-key")
}

func (ec *execContext) currentActiveMaps_autogen(olp, position lispObject) (lispObject, error) {
	return ec.stub("current-active-maps")
}

func (ec *execContext) keyBinding_autogen(key, accept_default, no_remap, position lispObject) (lispObject, error) {
	return ec.stub("key-binding")
}

func (ec *execContext) minorModeKeyBinding_autogen(key, accept_default lispObject) (lispObject, error) {
	return ec.stub("minor-mode-key-binding")
}

func (ec *execContext) useGlobalMap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("use-global-map")
}

func (ec *execContext) useLocalMap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("use-local-map")
}

func (ec *execContext) currentLocalMap_autogen() (lispObject, error) {
	return ec.stub("current-local-map")
}

func (ec *execContext) currentGlobalMap_autogen() (lispObject, error) {
	return ec.stub("current-global-map")
}

func (ec *execContext) currentMinorModeMaps_autogen() (lispObject, error) {
	return ec.stub("current-minor-mode-maps")
}

func (ec *execContext) accessibleKeymaps_autogen(keymap, prefix lispObject) (lispObject, error) {
	return ec.stub("accessible-keymaps")
}

func (ec *execContext) keyDescription_autogen(keys, prefix lispObject) (lispObject, error) {
	return ec.stub("key-description")
}

func (ec *execContext) singleKeyDescription_autogen(key, no_angles lispObject) (lispObject, error) {
	return ec.stub("single-key-description")
}

func (ec *execContext) textCharDescription_autogen(character lispObject) (lispObject, error) {
	return ec.stub("text-char-description")
}

func (ec *execContext) whereIsInternal_autogen(definition, keymap, firstonly, noindirect, no_remap lispObject) (lispObject, error) {
	return ec.stub("where-is-internal")
}

func (ec *execContext) describeBufferBindings_autogen(buffer, prefix, menus lispObject) (lispObject, error) {
	return ec.stub("describe-buffer-bindings")
}

func (ec *execContext) describeVector_autogen(vector, describer lispObject) (lispObject, error) {
	return ec.stub("describe-vector")
}

func (ec *execContext) helpDescribeVector_autogen(vector, prefix, describer, partial, shadow, entire_map, mention_shadow lispObject) (lispObject, error) {
	return ec.stub("help--describe-vector")
}

func (ec *execContext) xExportFrames_autogen(frames, type_ lispObject) (lispObject, error) {
	return ec.stub("x-export-frames")
}

func (ec *execContext) xExportFrames_1_autogen(frames, type_ lispObject) (lispObject, error) {
	return ec.stub("x-export-frames")
}

func (ec *execContext) pgtkSetMonitorScaleFactor_autogen(monitor_model, scale_factor lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-monitor-scale-factor")
}

func (ec *execContext) pgtkFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-restack")
}

func (ec *execContext) pgtkSetResource_autogen(attribute, value lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-resource")
}

func (ec *execContext) xServerMaxRequestSize_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size")
}

func (ec *execContext) xServerMaxRequestSize_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size")
}

func (ec *execContext) xServerMaxRequestSize_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size")
}

func (ec *execContext) pgtkFontName_autogen(name lispObject) (lispObject, error) {
	return ec.stub("pgtk-font-name")
}

func (ec *execContext) pgtkDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-display-monitor-attributes-list")
}

func (ec *execContext) pgtkFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-geometry")
}

func (ec *execContext) pgtkFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-edges")
}

func (ec *execContext) pgtkSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-mouse-absolute-pixel-position")
}

func (ec *execContext) pgtkMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("pgtk-mouse-absolute-pixel-position")
}

func (ec *execContext) pgtkPageSetupDialog_autogen() (lispObject, error) {
	return ec.stub("pgtk-page-setup-dialog")
}

func (ec *execContext) pgtkGetPageSetup_autogen() (lispObject, error) {
	return ec.stub("pgtk-get-page-setup")
}

func (ec *execContext) pgtkPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return ec.stub("pgtk-print-frames-dialog")
}

func (ec *execContext) xFileDialog_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog")
}

func (ec *execContext) xFileDialog_1_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog")
}

func (ec *execContext) xFileDialog_2_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog")
}

func (ec *execContext) xFileDialog_3_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog")
}

func (ec *execContext) pgtkBackendDisplayClass_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-backend-display-class")
}

func (ec *execContext) xGtkDebug_autogen(enable lispObject) (lispObject, error) {
	return ec.stub("x-gtk-debug")
}

func (ec *execContext) xGtkDebug_1_autogen(enable lispObject) (lispObject, error) {
	return ec.stub("x-gtk-debug")
}

func (ec *execContext) currentColumn_autogen() (lispObject, error) {
	return ec.stub("current-column")
}

func (ec *execContext) indentTo_autogen(column, minimum lispObject) (lispObject, error) {
	return ec.stub("indent-to")
}

func (ec *execContext) currentIndentation_autogen() (lispObject, error) {
	return ec.stub("current-indentation")
}

func (ec *execContext) moveToColumn_autogen(column, force lispObject) (lispObject, error) {
	return ec.stub("move-to-column")
}

func (ec *execContext) computeMotion_autogen(from, frompos, to, topos, width, offsets, window lispObject) (lispObject, error) {
	return ec.stub("compute-motion")
}

func (ec *execContext) lineNumberDisplayWidth_autogen(pixelwise lispObject) (lispObject, error) {
	return ec.stub("line-number-display-width")
}

func (ec *execContext) verticalMotion_autogen(lines, window, cur_col lispObject) (lispObject, error) {
	return ec.stub("vertical-motion")
}

func (ec *execContext) w32HasWinsock_autogen(load_now lispObject) (lispObject, error) {
	return ec.stub("w32-has-winsock")
}

func (ec *execContext) w32UnloadWinsock_autogen() (lispObject, error) {
	return ec.stub("w32-unload-winsock")
}

func (ec *execContext) w32ShortFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("w32-short-file-name")
}

func (ec *execContext) w32LongFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("w32-long-file-name")
}

func (ec *execContext) w32SetProcessPriority_autogen(process, priority lispObject) (lispObject, error) {
	return ec.stub("w32-set-process-priority")
}

func (ec *execContext) w32ApplicationType_autogen(program lispObject) (lispObject, error) {
	return ec.stub("w32-application-type")
}

func (ec *execContext) w32GetLocaleInfo_autogen(lcid, longform lispObject) (lispObject, error) {
	return ec.stub("w32-get-locale-info")
}

func (ec *execContext) w32GetCurrentLocaleId_autogen() (lispObject, error) {
	return ec.stub("w32-get-current-locale-id")
}

func (ec *execContext) w32GetValidLocaleIds_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-locale-ids")
}

func (ec *execContext) w32GetDefaultLocaleId_autogen(userp lispObject) (lispObject, error) {
	return ec.stub("w32-get-default-locale-id")
}

func (ec *execContext) w32SetCurrentLocale_autogen(lcid lispObject) (lispObject, error) {
	return ec.stub("w32-set-current-locale")
}

func (ec *execContext) w32GetValidCodepages_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-codepages")
}

func (ec *execContext) w32GetConsoleCodepage_autogen() (lispObject, error) {
	return ec.stub("w32-get-console-codepage")
}

func (ec *execContext) w32SetConsoleCodepage_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-set-console-codepage")
}

func (ec *execContext) w32GetConsoleOutputCodepage_autogen() (lispObject, error) {
	return ec.stub("w32-get-console-output-codepage")
}

func (ec *execContext) w32SetConsoleOutputCodepage_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-set-console-output-codepage")
}

func (ec *execContext) w32GetCodepageCharset_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-get-codepage-charset")
}

func (ec *execContext) w32GetValidKeyboardLayouts_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-keyboard-layouts")
}

func (ec *execContext) w32GetKeyboardLayout_autogen() (lispObject, error) {
	return ec.stub("w32-get-keyboard-layout")
}

func (ec *execContext) w32SetKeyboardLayout_autogen(layout lispObject) (lispObject, error) {
	return ec.stub("w32-set-keyboard-layout")
}

func (ec *execContext) windowp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("windowp")
}

func (ec *execContext) windowValidP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-valid-p")
}

func (ec *execContext) windowLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-live-p")
}

func (ec *execContext) windowFrame_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-frame")
}

func (ec *execContext) frameRootWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-root-window")
}

func (ec *execContext) minibufferWindow_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("minibuffer-window")
}

func (ec *execContext) windowMinibufferP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-minibuffer-p")
}

func (ec *execContext) frameFirstWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-first-window")
}

func (ec *execContext) frameSelectedWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-selected-window")
}

func (ec *execContext) frameOldSelectedWindow_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-old-selected-window")
}

func (ec *execContext) setFrameSelectedWindow_autogen(frame, window, norecord lispObject) (lispObject, error) {
	return ec.stub("set-frame-selected-window")
}

func (ec *execContext) selectedWindow_autogen() (lispObject, error) {
	return ec.stub("selected-window")
}

func (ec *execContext) oldSelectedWindow_autogen() (lispObject, error) {
	return ec.stub("old-selected-window")
}

func (ec *execContext) selectWindow_autogen(window, norecord lispObject) (lispObject, error) {
	return ec.stub("select-window")
}

func (ec *execContext) windowBuffer_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-buffer")
}

func (ec *execContext) windowOldBuffer_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-buffer")
}

func (ec *execContext) windowParent_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-parent")
}

func (ec *execContext) windowTopChild_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-top-child")
}

func (ec *execContext) windowLeftChild_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-left-child")
}

func (ec *execContext) windowNextSibling_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-next-sibling")
}

func (ec *execContext) windowPrevSibling_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-prev-sibling")
}

func (ec *execContext) windowCombinationLimit_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-combination-limit")
}

func (ec *execContext) setWindowCombinationLimit_autogen(window, limit lispObject) (lispObject, error) {
	return ec.stub("set-window-combination-limit")
}

func (ec *execContext) windowUseTime_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-use-time")
}

func (ec *execContext) windowBumpUseTime_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-bump-use-time")
}

func (ec *execContext) windowPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-width")
}

func (ec *execContext) windowPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-height")
}

func (ec *execContext) windowOldPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-pixel-width")
}

func (ec *execContext) windowOldPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-pixel-height")
}

func (ec *execContext) windowTotalHeight_autogen(window, round lispObject) (lispObject, error) {
	return ec.stub("window-total-height")
}

func (ec *execContext) windowTotalWidth_autogen(window, round lispObject) (lispObject, error) {
	return ec.stub("window-total-width")
}

func (ec *execContext) windowNewTotal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-total")
}

func (ec *execContext) windowNormalSize_autogen(window, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-normal-size")
}

func (ec *execContext) windowNewNormal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-normal")
}

func (ec *execContext) windowNewPixel_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-pixel")
}

func (ec *execContext) windowPixelLeft_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-left")
}

func (ec *execContext) windowPixelTop_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-top")
}

func (ec *execContext) windowLeftColumn_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-left-column")
}

func (ec *execContext) windowTopLine_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-top-line")
}

func (ec *execContext) windowBodyWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-body-width")
}

func (ec *execContext) windowBodyHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-body-height")
}

func (ec *execContext) windowOldBodyPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-body-pixel-width")
}

func (ec *execContext) windowOldBodyPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-body-pixel-height")
}

func (ec *execContext) windowModeLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-mode-line-height")
}

func (ec *execContext) windowHeaderLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-header-line-height")
}

func (ec *execContext) windowTabLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-tab-line-height")
}

func (ec *execContext) windowRightDividerWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-right-divider-width")
}

func (ec *execContext) windowBottomDividerWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-bottom-divider-width")
}

func (ec *execContext) windowScrollBarWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bar-width")
}

func (ec *execContext) windowScrollBarHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bar-height")
}

func (ec *execContext) windowHscroll_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-hscroll")
}

func (ec *execContext) setWindowHscroll_autogen(window, ncol lispObject) (lispObject, error) {
	return ec.stub("set-window-hscroll")
}

func (ec *execContext) coordinatesInWindowP_autogen(coordinates, window lispObject) (lispObject, error) {
	return ec.stub("coordinates-in-window-p")
}

func (ec *execContext) windowAt_autogen(x, y, frame lispObject) (lispObject, error) {
	return ec.stub("window-at")
}

func (ec *execContext) windowPoint_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-point")
}

func (ec *execContext) windowOldPoint_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-point")
}

func (ec *execContext) windowStart_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-start")
}

func (ec *execContext) windowEnd_autogen(window, update lispObject) (lispObject, error) {
	return ec.stub("window-end")
}

func (ec *execContext) setWindowPoint_autogen(window, pos lispObject) (lispObject, error) {
	return ec.stub("set-window-point")
}

func (ec *execContext) setWindowStart_autogen(window, pos, noforce lispObject) (lispObject, error) {
	return ec.stub("set-window-start")
}

func (ec *execContext) posVisibleInWindowP_autogen(pos, window, partially lispObject) (lispObject, error) {
	return ec.stub("pos-visible-in-window-p")
}

func (ec *execContext) windowLineHeight_autogen(line, window lispObject) (lispObject, error) {
	return ec.stub("window-line-height")
}

func (ec *execContext) windowLinesPixelDimensions_autogen(window, first, last, body, inverse, left lispObject) (lispObject, error) {
	return ec.stub("window-lines-pixel-dimensions")
}

func (ec *execContext) windowDedicatedP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-dedicated-p")
}

func (ec *execContext) setWindowDedicatedP_autogen(window, flag lispObject) (lispObject, error) {
	return ec.stub("set-window-dedicated-p")
}

func (ec *execContext) windowPrevBuffers_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-prev-buffers")
}

func (ec *execContext) setWindowPrevBuffers_autogen(window, prev_buffers lispObject) (lispObject, error) {
	return ec.stub("set-window-prev-buffers")
}

func (ec *execContext) windowNextBuffers_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-next-buffers")
}

func (ec *execContext) setWindowNextBuffers_autogen(window, next_buffers lispObject) (lispObject, error) {
	return ec.stub("set-window-next-buffers")
}

func (ec *execContext) windowParameters_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-parameters")
}

func (ec *execContext) windowParameter_autogen(window, parameter lispObject) (lispObject, error) {
	return ec.stub("window-parameter")
}

func (ec *execContext) setWindowParameter_autogen(window, parameter, value lispObject) (lispObject, error) {
	return ec.stub("set-window-parameter")
}

func (ec *execContext) windowDisplayTable_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-display-table")
}

func (ec *execContext) setWindowDisplayTable_autogen(window, table lispObject) (lispObject, error) {
	return ec.stub("set-window-display-table")
}

func (ec *execContext) nextWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("next-window")
}

func (ec *execContext) previousWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("previous-window")
}

func (ec *execContext) windowList_autogen(frame, minibuf, window lispObject) (lispObject, error) {
	return ec.stub("window-list")
}

func (ec *execContext) windowList1_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("window-list-1")
}

func (ec *execContext) getBufferWindow_autogen(buffer_or_name, all_frames lispObject) (lispObject, error) {
	return ec.stub("get-buffer-window")
}

func (ec *execContext) deleteOtherWindowsInternal_autogen(window, root lispObject) (lispObject, error) {
	return ec.stub("delete-other-windows-internal")
}

func (ec *execContext) runWindowConfigurationChangeHook_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("run-window-configuration-change-hook")
}

func (ec *execContext) runWindowScrollFunctions_autogen(window lispObject) (lispObject, error) {
	return ec.stub("run-window-scroll-functions")
}

func (ec *execContext) setWindowBuffer_autogen(window, buffer_or_name, keep_margins lispObject) (lispObject, error) {
	return ec.stub("set-window-buffer")
}

func (ec *execContext) forceWindowUpdate_autogen(object lispObject) (lispObject, error) {
	return ec.stub("force-window-update")
}

func (ec *execContext) setWindowNewPixel_autogen(window, size, add lispObject) (lispObject, error) {
	return ec.stub("set-window-new-pixel")
}

func (ec *execContext) setWindowNewTotal_autogen(window, size, add lispObject) (lispObject, error) {
	return ec.stub("set-window-new-total")
}

func (ec *execContext) setWindowNewNormal_autogen(window, size lispObject) (lispObject, error) {
	return ec.stub("set-window-new-normal")
}

func (ec *execContext) windowResizeApply_autogen(frame, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-resize-apply")
}

func (ec *execContext) windowResizeApplyTotal_autogen(frame, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-resize-apply-total")
}

func (ec *execContext) splitWindowInternal_autogen(old, pixel_size, side, normal_size lispObject) (lispObject, error) {
	return ec.stub("split-window-internal")
}

func (ec *execContext) deleteWindowInternal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("delete-window-internal")
}

func (ec *execContext) resizeMiniWindowInternal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("resize-mini-window-internal")
}

func (ec *execContext) scrollUp_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("scroll-up")
}

func (ec *execContext) scrollDown_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("scroll-down")
}

func (ec *execContext) otherWindowForScrolling_autogen() (lispObject, error) {
	return ec.stub("other-window-for-scrolling")
}

func (ec *execContext) scrollLeft_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return ec.stub("scroll-left")
}

func (ec *execContext) scrollRight_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return ec.stub("scroll-right")
}

func (ec *execContext) minibufferSelectedWindow_autogen() (lispObject, error) {
	return ec.stub("minibuffer-selected-window")
}

func (ec *execContext) recenter_autogen(arg, redisplay lispObject) (lispObject, error) {
	return ec.stub("recenter")
}

func (ec *execContext) windowTextWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-text-width")
}

func (ec *execContext) windowTextHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-text-height")
}

func (ec *execContext) moveToWindowLine_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("move-to-window-line")
}

func (ec *execContext) windowConfigurationP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-configuration-p")
}

func (ec *execContext) windowConfigurationFrame_autogen(config lispObject) (lispObject, error) {
	return ec.stub("window-configuration-frame")
}

func (ec *execContext) setWindowConfiguration_autogen(configuration, dont_set_frame, dont_set_miniwindow lispObject) (lispObject, error) {
	return ec.stub("set-window-configuration")
}

func (ec *execContext) currentWindowConfiguration_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("current-window-configuration")
}

func (ec *execContext) setWindowMargins_autogen(window, left_width, right_width lispObject) (lispObject, error) {
	return ec.stub("set-window-margins")
}

func (ec *execContext) windowMargins_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-margins")
}

func (ec *execContext) setWindowFringes_autogen(window, left_width, right_width, outside_margins, persistent lispObject) (lispObject, error) {
	return ec.stub("set-window-fringes")
}

func (ec *execContext) windowFringes_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-fringes")
}

func (ec *execContext) setWindowScrollBars_autogen(window, width, vertical_type, height, horizontal_type, persistent lispObject) (lispObject, error) {
	return ec.stub("set-window-scroll-bars")
}

func (ec *execContext) windowScrollBars_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bars")
}

func (ec *execContext) windowVscroll_autogen(window, pixels_p lispObject) (lispObject, error) {
	return ec.stub("window-vscroll")
}

func (ec *execContext) setWindowVscroll_autogen(window, vscroll, pixels_p, preserve_vscroll_p lispObject) (lispObject, error) {
	return ec.stub("set-window-vscroll")
}

func (ec *execContext) windowConfigurationEqualP_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("window-configuration-equal-p")
}

func (ec *execContext) characterp_autogen(object, ignore lispObject) (lispObject, error) {
	return ec.stub("characterp")
}

func (ec *execContext) maxChar_autogen(unicode lispObject) (lispObject, error) {
	return ec.stub("max-char")
}

func (ec *execContext) unibyteCharToMultibyte_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("unibyte-char-to-multibyte")
}

func (ec *execContext) multibyteCharToUnibyte_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("multibyte-char-to-unibyte")
}

func (ec *execContext) charWidth_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("char-width")
}

func (ec *execContext) stringWidth_autogen(str, from, to lispObject) (lispObject, error) {
	return ec.stub("string-width")
}

func (ec *execContext) string_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("string")
}

func (ec *execContext) unibyteString_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("unibyte-string")
}

func (ec *execContext) charResolveModifiers_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-resolve-modifiers")
}

func (ec *execContext) getByte_autogen(position, string lispObject) (lispObject, error) {
	return ec.stub("get-byte")
}

func (ec *execContext) charsetp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("charsetp")
}

func (ec *execContext) mapCharsetChars_autogen(function, charset, arg, from_code, to_code lispObject) (lispObject, error) {
	return ec.stub("map-charset-chars")
}

func (ec *execContext) defineCharsetInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("define-charset-internal")
}

func (ec *execContext) defineCharsetAlias_autogen(alias, charset lispObject) (lispObject, error) {
	return ec.stub("define-charset-alias")
}

func (ec *execContext) charsetPlist_autogen(charset lispObject) (lispObject, error) {
	return ec.stub("charset-plist")
}

func (ec *execContext) setCharsetPlist_autogen(charset, plist lispObject) (lispObject, error) {
	return ec.stub("set-charset-plist")
}

func (ec *execContext) unifyCharset_autogen(charset, unify_map, deunify lispObject) (lispObject, error) {
	return ec.stub("unify-charset")
}

func (ec *execContext) getUnusedIsoFinalChar_autogen(dimension, chars lispObject) (lispObject, error) {
	return ec.stub("get-unused-iso-final-char")
}

func (ec *execContext) declareEquivCharset_autogen(dimension, chars, final_char, charset lispObject) (lispObject, error) {
	return ec.stub("declare-equiv-charset")
}

func (ec *execContext) findCharsetRegion_autogen(beg, end, table lispObject) (lispObject, error) {
	return ec.stub("find-charset-region")
}

func (ec *execContext) findCharsetString_autogen(str, table lispObject) (lispObject, error) {
	return ec.stub("find-charset-string")
}

func (ec *execContext) decodeChar_autogen(charset, code_point lispObject) (lispObject, error) {
	return ec.stub("decode-char")
}

func (ec *execContext) encodeChar_autogen(ch, charset lispObject) (lispObject, error) {
	return ec.stub("encode-char")
}

func (ec *execContext) makeChar_autogen(charset, code1, code2, code3, code4 lispObject) (lispObject, error) {
	return ec.stub("make-char")
}

func (ec *execContext) splitChar_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("split-char")
}

func (ec *execContext) charCharset_autogen(ch, restriction lispObject) (lispObject, error) {
	return ec.stub("char-charset")
}

func (ec *execContext) charsetAfter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("charset-after")
}

func (ec *execContext) isoCharset_autogen(dimension, chars, final_char lispObject) (lispObject, error) {
	return ec.stub("iso-charset")
}

func (ec *execContext) clearCharsetMaps_autogen() (lispObject, error) {
	return ec.stub("clear-charset-maps")
}

func (ec *execContext) charsetPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return ec.stub("charset-priority-list")
}

func (ec *execContext) setCharsetPriority_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("set-charset-priority")
}

func (ec *execContext) charsetIdInternal_autogen(charset lispObject) (lispObject, error) {
	return ec.stub("charset-id-internal")
}

func (ec *execContext) sortCharsets_autogen(charsets lispObject) (lispObject, error) {
	return ec.stub("sort-charsets")
}

func (ec *execContext) pgtkOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return ec.stub("pgtk-own-selection-internal")
}

func (ec *execContext) pgtkGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-get-selection-internal")
}

func (ec *execContext) pgtkDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-disown-selection-internal")
}

func (ec *execContext) pgtkSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-selection-owner-p")
}

func (ec *execContext) pgtkSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-selection-exists-p")
}

func (ec *execContext) pgtkRegisterDndTargets_autogen(frame, targets lispObject) (lispObject, error) {
	return ec.stub("pgtk-register-dnd-targets")
}

func (ec *execContext) pgtkDropFinish_autogen(success, timestamp, delete lispObject) (lispObject, error) {
	return ec.stub("pgtk-drop-finish")
}

func (ec *execContext) pgtkUpdateDropStatus_autogen(action, timestamp lispObject) (lispObject, error) {
	return ec.stub("pgtk-update-drop-status")
}

func (ec *execContext) lookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("looking-at")
}

func (ec *execContext) posixLookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("posix-looking-at")
}

func (ec *execContext) stringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("string-match")
}

func (ec *execContext) posixStringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("posix-string-match")
}

func (ec *execContext) searchBackward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("search-backward")
}

func (ec *execContext) searchForward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("search-forward")
}

func (ec *execContext) reSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("re-search-backward")
}

func (ec *execContext) reSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("re-search-forward")
}

func (ec *execContext) posixSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("posix-search-backward")
}

func (ec *execContext) posixSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("posix-search-forward")
}

func (ec *execContext) replaceMatch_autogen(newtext, fixedcase, literal, string, subexp lispObject) (lispObject, error) {
	return ec.stub("replace-match")
}

func (ec *execContext) matchBeginning_autogen(subexp lispObject) (lispObject, error) {
	return ec.stub("match-beginning")
}

func (ec *execContext) matchEnd_autogen(subexp lispObject) (lispObject, error) {
	return ec.stub("match-end")
}

func (ec *execContext) matchData_autogen(integers, reuse, reseat lispObject) (lispObject, error) {
	return ec.stub("match-data")
}

func (ec *execContext) setMatchData_autogen(list, reseat lispObject) (lispObject, error) {
	return ec.stub("set-match-data")
}

func (ec *execContext) matchDataTranslate_autogen(n lispObject) (lispObject, error) {
	return ec.stub("match-data--translate")
}

func (ec *execContext) regexpQuote_autogen(string lispObject) (lispObject, error) {
	return ec.stub("regexp-quote")
}

func (ec *execContext) newlineCacheCheck_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("newline-cache-check")
}

func (ec *execContext) deleteTerminal_autogen(terminal, force lispObject) (lispObject, error) {
	return ec.stub("delete-terminal")
}

func (ec *execContext) frameTerminal_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-terminal")
}

func (ec *execContext) terminalLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("terminal-live-p")
}

func (ec *execContext) terminalList_autogen() (lispObject, error) {
	return ec.stub("terminal-list")
}

func (ec *execContext) terminalName_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-name")
}

func (ec *execContext) terminalParameters_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-parameters")
}

func (ec *execContext) terminalParameter_autogen(terminal, parameter lispObject) (lispObject, error) {
	return ec.stub("terminal-parameter")
}

func (ec *execContext) setTerminalParameter_autogen(terminal, parameter, value lispObject) (lispObject, error) {
	return ec.stub("set-terminal-parameter")
}

func (ec *execContext) menuBarMenuAtXY_autogen(x, y, frame lispObject) (lispObject, error) {
	return ec.stub("menu-bar-menu-at-x-y")
}

func (ec *execContext) xPopupMenu_autogen(position, menu lispObject) (lispObject, error) {
	return ec.stub("x-popup-menu")
}

func (ec *execContext) xPopupDialog_autogen(position, contents, header lispObject) (lispObject, error) {
	return ec.stub("x-popup-dialog")
}

func (ec *execContext) markerBuffer_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-buffer")
}

func (ec *execContext) markerPosition_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-position")
}

func (ec *execContext) setMarker_autogen(marker, position, buffer lispObject) (lispObject, error) {
	return ec.stub("set-marker")
}

func (ec *execContext) copyMarker_autogen(marker, type_ lispObject) (lispObject, error) {
	return ec.stub("copy-marker")
}

func (ec *execContext) markerInsertionType_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-insertion-type")
}

func (ec *execContext) setMarkerInsertionType_autogen(marker, type_ lispObject) (lispObject, error) {
	return ec.stub("set-marker-insertion-type")
}

func (ec *execContext) w16SetClipboardData_autogen(string, frame lispObject) (lispObject, error) {
	return ec.stub("w16-set-clipboard-data")
}

func (ec *execContext) w16GetClipboardData_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("w16-get-clipboard-data")
}

func (ec *execContext) w16SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w16-selection-exists-p")
}

func (ec *execContext) dumpRedisplayHistory_autogen() (lispObject, error) {
	return ec.stub("dump-redisplay-history")
}

func (ec *execContext) redrawFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("redraw-frame")
}

func (ec *execContext) redrawDisplay_autogen() (lispObject, error) {
	return ec.stub("redraw-display")
}

func (ec *execContext) displayUpdateForMouseMovement_autogen(mouse_x, mouse_y lispObject) (lispObject, error) {
	return ec.stub("display--update-for-mouse-movement")
}

func (ec *execContext) openTermscript_autogen(file lispObject) (lispObject, error) {
	return ec.stub("open-termscript")
}

func (ec *execContext) sendStringToTerminal_autogen(string, terminal lispObject) (lispObject, error) {
	return ec.stub("send-string-to-terminal")
}

func (ec *execContext) ding_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ding")
}

func (ec *execContext) sleepFor_autogen(seconds, milliseconds lispObject) (lispObject, error) {
	return ec.stub("sleep-for")
}

func (ec *execContext) redisplay_autogen(force lispObject) (lispObject, error) {
	return ec.stub("redisplay")
}

func (ec *execContext) frameOrBufferChangedP_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("frame-or-buffer-changed-p")
}

func (ec *execContext) internalShowCursor_autogen(window, show lispObject) (lispObject, error) {
	return ec.stub("internal-show-cursor")
}

func (ec *execContext) internalShowCursorP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("internal-show-cursor-p")
}

func (ec *execContext) directoryFiles_autogen(directory, full, match, nosort, count lispObject) (lispObject, error) {
	return ec.stub("directory-files")
}

func (ec *execContext) directoryFilesAndAttributes_autogen(directory, full, match, nosort, id_format, count lispObject) (lispObject, error) {
	return ec.stub("directory-files-and-attributes")
}

func (ec *execContext) fileNameCompletion_autogen(file, directory, predicate lispObject) (lispObject, error) {
	return ec.stub("file-name-completion")
}

func (ec *execContext) fileNameAllCompletions_autogen(file, directory lispObject) (lispObject, error) {
	return ec.stub("file-name-all-completions")
}

func (ec *execContext) fileAttributes_autogen(filename, id_format lispObject) (lispObject, error) {
	return ec.stub("file-attributes")
}

func (ec *execContext) fileAttributesLessp_autogen(f1, f2 lispObject) (lispObject, error) {
	return ec.stub("file-attributes-lessp")
}

func (ec *execContext) systemUsers_autogen() (lispObject, error) {
	return ec.stub("system-users")
}

func (ec *execContext) systemGroups_autogen() (lispObject, error) {
	return ec.stub("system-groups")
}

func (ec *execContext) debugTimerCheck_autogen() (lispObject, error) {
	return ec.stub("debug-timer-check")
}

func (ec *execContext) upcase_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("upcase")
}

func (ec *execContext) downcase_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("downcase")
}

func (ec *execContext) capitalize_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("capitalize")
}

func (ec *execContext) upcaseInitials_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("upcase-initials")
}

func (ec *execContext) upcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("upcase-region")
}

func (ec *execContext) downcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("downcase-region")
}

func (ec *execContext) capitalizeRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("capitalize-region")
}

func (ec *execContext) upcaseInitialsRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("upcase-initials-region")
}

func (ec *execContext) upcaseWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("upcase-word")
}

func (ec *execContext) downcaseWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("downcase-word")
}

func (ec *execContext) capitalizeWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("capitalize-word")
}

func (ec *execContext) xWmSetSizeHint_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-wm-set-size-hint")
}

func (ec *execContext) xServerInputExtensionVersion_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-input-extension-version")
}

func (ec *execContext) xDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-monitor-attributes-list")
}

func (ec *execContext) xFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-frame-geometry")
}

func (ec *execContext) xFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("x-frame-edges")
}

func (ec *execContext) xFrameListZOrder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-frame-list-z-order")
}

func (ec *execContext) xFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("x-frame-restack")
}

func (ec *execContext) xMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("x-mouse-absolute-pixel-position")
}

func (ec *execContext) xSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("x-set-mouse-absolute-pixel-position")
}

func (ec *execContext) xBeginDrag_autogen(targets, action, frame, return_frame, allow_current_frame, follow_tooltip lispObject) (lispObject, error) {
	return ec.stub("x-begin-drag")
}

func (ec *execContext) xSynchronize_autogen(on, terminal lispObject) (lispObject, error) {
	return ec.stub("x-synchronize")
}

func (ec *execContext) xSynchronize_1_autogen(on, display lispObject) (lispObject, error) {
	return ec.stub("x-synchronize")
}

func (ec *execContext) xChangeWindowProperty_autogen(prop, value, frame, type_, format, outer_p, window_id lispObject) (lispObject, error) {
	return ec.stub("x-change-window-property")
}

func (ec *execContext) xChangeWindowProperty_1_autogen(prop, value, frame, type_, format, outer_p lispObject) (lispObject, error) {
	return ec.stub("x-change-window-property")
}

func (ec *execContext) xDeleteWindowProperty_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return ec.stub("x-delete-window-property")
}

func (ec *execContext) xDeleteWindowProperty_1_autogen(prop, frame lispObject) (lispObject, error) {
	return ec.stub("x-delete-window-property")
}

func (ec *execContext) xWindowProperty_autogen(prop, frame, type_, window_id, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return ec.stub("x-window-property")
}

func (ec *execContext) xWindowProperty_1_autogen(prop, frame, type_, source, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return ec.stub("x-window-property")
}

func (ec *execContext) xWindowPropertyAttributes_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return ec.stub("x-window-property-attributes")
}

func (ec *execContext) xTranslateCoordinates_autogen(frame, source_window, dest_window, source_x, source_y, require_child lispObject) (lispObject, error) {
	return ec.stub("x-translate-coordinates")
}

func (ec *execContext) xUsesOldGtkDialog_autogen() (lispObject, error) {
	return ec.stub("x-uses-old-gtk-dialog")
}

func (ec *execContext) xBackspaceDeleteKeysP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-backspace-delete-keys-p")
}

func (ec *execContext) xGetModifierMasks_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-get-modifier-masks")
}

func (ec *execContext) xPageSetupDialog_autogen() (lispObject, error) {
	return ec.stub("x-page-setup-dialog")
}

func (ec *execContext) xGetPageSetup_autogen() (lispObject, error) {
	return ec.stub("x-get-page-setup")
}

func (ec *execContext) xPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return ec.stub("x-print-frames-dialog")
}

func (ec *execContext) xDisplayLastUserTime_autogen(time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-set-last-user-time")
}

func (ec *execContext) xInternalFocusInputContext_autogen(focus lispObject) (lispObject, error) {
	return ec.stub("x-internal-focus-input-context")
}

func (ec *execContext) setScreenColor_autogen(foreground, background lispObject) (lispObject, error) {
	return ec.stub("set-screen-color")
}

func (ec *execContext) getScreenColor_autogen() (lispObject, error) {
	return ec.stub("get-screen-color")
}

func (ec *execContext) setCursorSize_autogen(size lispObject) (lispObject, error) {
	return ec.stub("set-cursor-size")
}

func (ec *execContext) caseTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("case-table-p")
}

func (ec *execContext) currentCaseTable_autogen() (lispObject, error) {
	return ec.stub("current-case-table")
}

func (ec *execContext) standardCaseTable_autogen() (lispObject, error) {
	return ec.stub("standard-case-table")
}

func (ec *execContext) setCaseTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-case-table")
}

func (ec *execContext) setStandardCaseTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-standard-case-table")
}

func (ec *execContext) haikuSelectionTimestamp_autogen(clipboard lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-timestamp")
}

func (ec *execContext) haikuSelectionData_autogen(clipboard, name lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-data")
}

func (ec *execContext) haikuSelectionPut_autogen(clipboard, name, data, clear lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-put")
}

func (ec *execContext) haikuSelectionOwnerP_autogen(selection lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-owner-p")
}

func (ec *execContext) haikuDragMessage_autogen(frame, message, allow_same_frame, follow_tooltip lispObject) (lispObject, error) {
	return ec.stub("haiku-drag-message")
}

func (ec *execContext) haikuRosterLaunch_autogen(file_or_type, args lispObject) (lispObject, error) {
	return ec.stub("haiku-roster-launch")
}

func (ec *execContext) haikuWriteNodeAttribute_autogen(file, name, message lispObject) (lispObject, error) {
	return ec.stub("haiku-write-node-attribute")
}

func (ec *execContext) haikuSendMessage_autogen(program, message lispObject) (lispObject, error) {
	return ec.stub("haiku-send-message")
}

func (ec *execContext) makeXwidget_autogen(type_, title, width, height, arguments, buffer, related lispObject) (lispObject, error) {
	return ec.stub("make-xwidget")
}

func (ec *execContext) xwidgetLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidget-live-p")
}

func (ec *execContext) xwidgetPerformLispyEvent_autogen(xwidget, event, frame lispObject) (lispObject, error) {
	return ec.stub("xwidget-perform-lispy-event")
}

func (ec *execContext) getBufferXwidgets_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("get-buffer-xwidgets")
}

func (ec *execContext) xwidgetWebkitUri_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-uri")
}

func (ec *execContext) xwidgetWebkitTitle_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-title")
}

func (ec *execContext) xwidgetWebkitGotoUri_autogen(xwidget, uri lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-goto-uri")
}

func (ec *execContext) xwidgetWebkitGotoHistory_autogen(xwidget, rel_pos lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-goto-history")
}

func (ec *execContext) xwidgetWebkitZoom_autogen(xwidget, factor lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-zoom")
}

func (ec *execContext) xwidgetWebkitExecuteScript_autogen(xwidget, script, fun lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-execute-script")
}

func (ec *execContext) xwidgetResize_autogen(xwidget, new_width, new_height lispObject) (lispObject, error) {
	return ec.stub("xwidget-resize")
}

func (ec *execContext) xwidgetSizeRequest_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-size-request")
}

func (ec *execContext) xwidgetp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidgetp")
}

func (ec *execContext) xwidgetViewP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-p")
}

func (ec *execContext) xwidgetInfo_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-info")
}

func (ec *execContext) xwidgetViewInfo_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-info")
}

func (ec *execContext) xwidgetViewModel_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-model")
}

func (ec *execContext) xwidgetViewWindow_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-window")
}

func (ec *execContext) deleteXwidgetView_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("delete-xwidget-view")
}

func (ec *execContext) xwidgetViewLookup_autogen(xwidget, window lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-lookup")
}

func (ec *execContext) xwidgetPlist_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-plist")
}

func (ec *execContext) xwidgetBuffer_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-buffer")
}

func (ec *execContext) setXwidgetBuffer_autogen(xwidget, buffer lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-buffer")
}

func (ec *execContext) setXwidgetPlist_autogen(xwidget, plist lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-plist")
}

func (ec *execContext) setXwidgetQueryOnExitFlag_autogen(xwidget, flag lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-query-on-exit-flag")
}

func (ec *execContext) xwidgetQueryOnExitFlag_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-query-on-exit-flag")
}

func (ec *execContext) xwidgetWebkitSearch_autogen(query, xwidget, case_insensitive, backwards, wrap_around lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-search")
}

func (ec *execContext) xwidgetWebkitNextResult_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-next-result")
}

func (ec *execContext) xwidgetWebkitPreviousResult_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-previous-result")
}

func (ec *execContext) xwidgetWebkitFinishSearch_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-finish-search")
}

func (ec *execContext) killXwidget_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("kill-xwidget")
}

func (ec *execContext) xwidgetWebkitLoadHtml_autogen(xwidget, text, base_uri lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-load-html")
}

func (ec *execContext) xwidgetWebkitBackForwardList_autogen(xwidget, limit lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-back-forward-list")
}

func (ec *execContext) xwidgetWebkitEstimatedLoadProgress_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-estimated-load-progress")
}

func (ec *execContext) xwidgetWebkitSetCookieStorageFile_autogen(xwidget, file lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-set-cookie-storage-file")
}

func (ec *execContext) xwidgetWebkitStopLoading_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-stop-loading")
}

func (ec *execContext) zlibAvailableP_autogen() (lispObject, error) {
	return ec.stub("zlib-available-p")
}

func (ec *execContext) zlibDecompressRegion_autogen(start, end, allow_partial lispObject) (lispObject, error) {
	return ec.stub("zlib-decompress-region")
}

func (ec *execContext) w32DefineRgbColor_autogen(red, green, blue, name lispObject) (lispObject, error) {
	return ec.stub("w32-define-rgb-color")
}

func (ec *execContext) w32DisplayMonitorAttributesList_autogen(display lispObject) (lispObject, error) {
	return ec.stub("w32-display-monitor-attributes-list")
}

func (ec *execContext) setMessageBeep_autogen(sound lispObject) (lispObject, error) {
	return ec.stub("set-message-beep")
}

func (ec *execContext) systemMoveFileToTrash_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("system-move-file-to-trash")
}

func (ec *execContext) w32SendSysCommand_autogen(command, frame lispObject) (lispObject, error) {
	return ec.stub("w32-send-sys-command")
}

func (ec *execContext) w32ShellExecute_autogen(operation, document, parameters, show_flag lispObject) (lispObject, error) {
	return ec.stub("w32-shell-execute")
}

func (ec *execContext) w32RegisterHotKey_autogen(key lispObject) (lispObject, error) {
	return ec.stub("w32-register-hot-key")
}

func (ec *execContext) w32UnregisterHotKey_autogen(key lispObject) (lispObject, error) {
	return ec.stub("w32-unregister-hot-key")
}

func (ec *execContext) w32RegisteredHotKeys_autogen() (lispObject, error) {
	return ec.stub("w32-registered-hot-keys")
}

func (ec *execContext) w32ReconstructHotKey_autogen(hotkeyid lispObject) (lispObject, error) {
	return ec.stub("w32-reconstruct-hot-key")
}

func (ec *execContext) w32ToggleLockKey_autogen(key, new_state lispObject) (lispObject, error) {
	return ec.stub("w32-toggle-lock-key")
}

func (ec *execContext) w32WindowExistsP_autogen(class, name lispObject) (lispObject, error) {
	return ec.stub("w32-window-exists-p")
}

func (ec *execContext) w32FrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("w32-frame-geometry")
}

func (ec *execContext) w32FrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("w32-frame-edges")
}

func (ec *execContext) w32FrameListZOrder_autogen(display lispObject) (lispObject, error) {
	return ec.stub("w32-frame-list-z-order")
}

func (ec *execContext) w32FrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("w32-frame-restack")
}

func (ec *execContext) w32MouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("w32-mouse-absolute-pixel-position")
}

func (ec *execContext) w32SetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("w32-set-mouse-absolute-pixel-position")
}

func (ec *execContext) defaultPrinterName_autogen() (lispObject, error) {
	return ec.stub("default-printer-name")
}

func (ec *execContext) w32MenuBarInUse_autogen() (lispObject, error) {
	return ec.stub("w32--menu-bar-in-use")
}

func (ec *execContext) w32NotificationNotify_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("w32-notification-notify")
}

func (ec *execContext) w32NotificationClose_autogen(id lispObject) (lispObject, error) {
	return ec.stub("w32-notification-close")
}

func (ec *execContext) w32GetImeOpenStatus_autogen() (lispObject, error) {
	return ec.stub("w32-get-ime-open-status")
}

func (ec *execContext) w32SetImeOpenStatus_autogen(status lispObject) (lispObject, error) {
	return ec.stub("w32-set-ime-open-status")
}

func (ec *execContext) w32ReadRegistry_autogen(root, key, name lispObject) (lispObject, error) {
	return ec.stub("w32-read-registry")
}

func (ec *execContext) w32SetWallpaper_autogen(image_file lispObject) (lispObject, error) {
	return ec.stub("w32-set-wallpaper")
}

func (ec *execContext) makeCategorySet_autogen(categories lispObject) (lispObject, error) {
	return ec.stub("make-category-set")
}

func (ec *execContext) defineCategory_autogen(category, docstring, table lispObject) (lispObject, error) {
	return ec.stub("define-category")
}

func (ec *execContext) categoryDocstring_autogen(category, table lispObject) (lispObject, error) {
	return ec.stub("category-docstring")
}

func (ec *execContext) getUnusedCategory_autogen(table lispObject) (lispObject, error) {
	return ec.stub("get-unused-category")
}

func (ec *execContext) categoryTableP_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("category-table-p")
}

func (ec *execContext) categoryTable_autogen() (lispObject, error) {
	return ec.stub("category-table")
}

func (ec *execContext) standardCategoryTable_autogen() (lispObject, error) {
	return ec.stub("standard-category-table")
}

func (ec *execContext) copyCategoryTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-category-table")
}

func (ec *execContext) makeCategoryTable_autogen() (lispObject, error) {
	return ec.stub("make-category-table")
}

func (ec *execContext) setCategoryTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-category-table")
}

func (ec *execContext) charCategorySet_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("char-category-set")
}

func (ec *execContext) categorySetMnemonics_autogen(category_set lispObject) (lispObject, error) {
	return ec.stub("category-set-mnemonics")
}

func (ec *execContext) modifyCategoryEntry_autogen(character, category, table, reset lispObject) (lispObject, error) {
	return ec.stub("modify-category-entry")
}

func (ec *execContext) jsonAvailableP_autogen() (lispObject, error) {
	return ec.stub("json--available-p")
}

func (ec *execContext) jsonSerialize_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-serialize")
}

func (ec *execContext) jsonInsert_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-insert")
}

func (ec *execContext) jsonParseString_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-parse-string")
}

func (ec *execContext) jsonParseBuffer_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-parse-buffer")
}

func (ec *execContext) interactive_autogen(args lispObject) (lispObject, error) {
	return ec.stub("interactive")
}

func (ec *execContext) funcallInteractively_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("funcall-interactively")
}

func (ec *execContext) callInteractively_autogen(function, record_flag, keys lispObject) (lispObject, error) {
	return ec.stub("call-interactively")
}

func (ec *execContext) prefixNumericValue_autogen(raw lispObject) (lispObject, error) {
	return ec.stub("prefix-numeric-value")
}

func (ec *execContext) processp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("processp")
}

func (ec *execContext) getProcess_autogen(name lispObject) (lispObject, error) {
	return ec.stub("get-process")
}

func (ec *execContext) deleteProcess_autogen(process lispObject) (lispObject, error) {
	return ec.stub("delete-process")
}

func (ec *execContext) processStatus_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-status")
}

func (ec *execContext) processExitStatus_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-exit-status")
}

func (ec *execContext) processId_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-id")
}

func (ec *execContext) processName_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-name")
}

func (ec *execContext) processCommand_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-command")
}

func (ec *execContext) processTtyName_autogen(process, stream lispObject) (lispObject, error) {
	return ec.stub("process-tty-name")
}

func (ec *execContext) setProcessBuffer_autogen(process, buffer lispObject) (lispObject, error) {
	return ec.stub("set-process-buffer")
}

func (ec *execContext) processBuffer_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-buffer")
}

func (ec *execContext) processMark_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-mark")
}

func (ec *execContext) setProcessFilter_autogen(process, filter lispObject) (lispObject, error) {
	return ec.stub("set-process-filter")
}

func (ec *execContext) processFilter_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-filter")
}

func (ec *execContext) setProcessSentinel_autogen(process, sentinel lispObject) (lispObject, error) {
	return ec.stub("set-process-sentinel")
}

func (ec *execContext) processSentinel_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-sentinel")
}

func (ec *execContext) setProcessThread_autogen(process, thread lispObject) (lispObject, error) {
	return ec.stub("set-process-thread")
}

func (ec *execContext) processThread_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-thread")
}

func (ec *execContext) setProcessWindowSize_autogen(process, height, width lispObject) (lispObject, error) {
	return ec.stub("set-process-window-size")
}

func (ec *execContext) setProcessInheritCodingSystemFlag_autogen(process, flag lispObject) (lispObject, error) {
	return ec.stub("set-process-inherit-coding-system-flag")
}

func (ec *execContext) setProcessQueryOnExitFlag_autogen(process, flag lispObject) (lispObject, error) {
	return ec.stub("set-process-query-on-exit-flag")
}

func (ec *execContext) processQueryOnExitFlag_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-query-on-exit-flag")
}

func (ec *execContext) processContact_autogen(process, key, no_block lispObject) (lispObject, error) {
	return ec.stub("process-contact")
}

func (ec *execContext) processPlist_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-plist")
}

func (ec *execContext) setProcessPlist_autogen(process, plist lispObject) (lispObject, error) {
	return ec.stub("set-process-plist")
}

func (ec *execContext) processConnection_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-connection")
}

func (ec *execContext) processType_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-type")
}

func (ec *execContext) formatNetworkAddress_autogen(address, omit_port lispObject) (lispObject, error) {
	return ec.stub("format-network-address")
}

func (ec *execContext) processList_autogen() (lispObject, error) {
	return ec.stub("process-list")
}

func (ec *execContext) makeProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-process")
}

func (ec *execContext) makePipeProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-pipe-process")
}

func (ec *execContext) processDatagramAddress_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-datagram-address")
}

func (ec *execContext) setProcessDatagramAddress_autogen(process, address lispObject) (lispObject, error) {
	return ec.stub("set-process-datagram-address")
}

func (ec *execContext) setNetworkProcessOption_autogen(process, option, value, no_error lispObject) (lispObject, error) {
	return ec.stub("set-network-process-option")
}

func (ec *execContext) serialProcessConfigure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("serial-process-configure")
}

func (ec *execContext) makeSerialProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-serial-process")
}

func (ec *execContext) makeNetworkProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-network-process")
}

func (ec *execContext) networkInterfaceList_autogen(full, family lispObject) (lispObject, error) {
	return ec.stub("network-interface-list")
}

func (ec *execContext) networkInterfaceInfo_autogen(ifname lispObject) (lispObject, error) {
	return ec.stub("network-interface-info")
}

func (ec *execContext) networkLookupAddressInfo_autogen(name, family, hint lispObject) (lispObject, error) {
	return ec.stub("network-lookup-address-info")
}

func (ec *execContext) acceptProcessOutput_autogen(process, seconds, millisec, just_this_one lispObject) (lispObject, error) {
	return ec.stub("accept-process-output")
}

func (ec *execContext) internalDefaultProcessFilter_autogen(proc, text lispObject) (lispObject, error) {
	return ec.stub("internal-default-process-filter")
}

func (ec *execContext) processSendRegion_autogen(process, start, end lispObject) (lispObject, error) {
	return ec.stub("process-send-region")
}

func (ec *execContext) processSendString_autogen(process, string lispObject) (lispObject, error) {
	return ec.stub("process-send-string")
}

func (ec *execContext) processRunningChildP_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-running-child-p")
}

func (ec *execContext) internalDefaultInterruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("internal-default-interrupt-process")
}

func (ec *execContext) interruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("interrupt-process")
}

func (ec *execContext) killProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("kill-process")
}

func (ec *execContext) quitProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("quit-process")
}

func (ec *execContext) stopProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("stop-process")
}

func (ec *execContext) continueProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("continue-process")
}

func (ec *execContext) internalDefaultSignalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return ec.stub("internal-default-signal-process")
}

func (ec *execContext) signalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return ec.stub("signal-process")
}

func (ec *execContext) processSendEof_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-send-eof")
}

func (ec *execContext) internalDefaultProcessSentinel_autogen(proc, msg lispObject) (lispObject, error) {
	return ec.stub("internal-default-process-sentinel")
}

func (ec *execContext) setProcessCodingSystem_autogen(process, decoding, encoding lispObject) (lispObject, error) {
	return ec.stub("set-process-coding-system")
}

func (ec *execContext) processCodingSystem_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-coding-system")
}

func (ec *execContext) getBufferProcess_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("get-buffer-process")
}

func (ec *execContext) processInheritCodingSystemFlag_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-inherit-coding-system-flag")
}

func (ec *execContext) waitingForUserInputP_autogen() (lispObject, error) {
	return ec.stub("waiting-for-user-input-p")
}

func (ec *execContext) listSystemProcesses_autogen() (lispObject, error) {
	return ec.stub("list-system-processes")
}

func (ec *execContext) processAttributes_autogen(pid lispObject) (lispObject, error) {
	return ec.stub("process-attributes")
}

func (ec *execContext) numProcessors_autogen(query lispObject) (lispObject, error) {
	return ec.stub("num-processors")
}

func (ec *execContext) signalNames_autogen() (lispObject, error) {
	return ec.stub("signal-names")
}

func (ec *execContext) makeString_autogen(length, init, multibyte lispObject) (lispObject, error) {
	return ec.stub("make-string")
}

func (ec *execContext) makeBoolVector_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-bool-vector")
}

func (ec *execContext) boolVector_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("bool-vector")
}

func (ec *execContext) cons_autogen(car, cdr lispObject) (lispObject, error) {
	return ec.stub("cons")
}

func (ec *execContext) list_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("list")
}

func (ec *execContext) makeList_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-list")
}

func (ec *execContext) makeRecord_autogen(type_, slots, init lispObject) (lispObject, error) {
	return ec.stub("make-record")
}

func (ec *execContext) record_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("record")
}

func (ec *execContext) makeVector_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-vector")
}

func (ec *execContext) vector_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("vector")
}

func (ec *execContext) makeByteCode_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-byte-code")
}

func (ec *execContext) makeClosure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-closure")
}

func (ec *execContext) makeSymbol_autogen(name lispObject) (lispObject, error) {
	return ec.stub("make-symbol")
}

func (ec *execContext) makeMarker_autogen() (lispObject, error) {
	return ec.stub("make-marker")
}

func (ec *execContext) makeFinalizer_autogen(function lispObject) (lispObject, error) {
	return ec.stub("make-finalizer")
}

func (ec *execContext) purecopy_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("purecopy")
}

func (ec *execContext) garbageCollect_autogen() (lispObject, error) {
	return ec.stub("garbage-collect")
}

func (ec *execContext) garbageCollectMaybe_autogen(factor lispObject) (lispObject, error) {
	return ec.stub("garbage-collect-maybe")
}

func (ec *execContext) memoryInfo_autogen() (lispObject, error) {
	return ec.stub("memory-info")
}

func (ec *execContext) memoryUseCounts_autogen() (lispObject, error) {
	return ec.stub("memory-use-counts")
}

func (ec *execContext) mallocInfo_autogen() (lispObject, error) {
	return ec.stub("malloc-info")
}

func (ec *execContext) mallocTrim_autogen(leave_padding lispObject) (lispObject, error) {
	return ec.stub("malloc-trim")
}

func (ec *execContext) suspiciousObject_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("suspicious-object")
}

func (ec *execContext) acos_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("acos")
}

func (ec *execContext) asin_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("asin")
}

func (ec *execContext) atan_autogen(y, x lispObject) (lispObject, error) {
	return ec.stub("atan")
}

func (ec *execContext) cos_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("cos")
}

func (ec *execContext) sin_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("sin")
}

func (ec *execContext) tan_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("tan")
}

func (ec *execContext) isnan_autogen(x lispObject) (lispObject, error) {
	return ec.stub("isnan")
}

func (ec *execContext) copysign_autogen(x1, x2 lispObject) (lispObject, error) {
	return ec.stub("copysign")
}

func (ec *execContext) frexp_autogen(x lispObject) (lispObject, error) {
	return ec.stub("frexp")
}

func (ec *execContext) ldexp_autogen(sgnfcand, exponent lispObject) (lispObject, error) {
	return ec.stub("ldexp")
}

func (ec *execContext) exp_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("exp")
}

func (ec *execContext) expt_autogen(arg1, arg2 lispObject) (lispObject, error) {
	return ec.stub("expt")
}

func (ec *execContext) log_autogen(arg, base lispObject) (lispObject, error) {
	return ec.stub("log")
}

func (ec *execContext) sqrt_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("sqrt")
}

func (ec *execContext) abs_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("abs")
}

func (ec *execContext) float_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("float")
}

func (ec *execContext) logb_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("logb")
}

func (ec *execContext) ceiling_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("ceiling")
}

func (ec *execContext) floor_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("floor")
}

func (ec *execContext) round_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("round")
}

func (ec *execContext) truncate_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("truncate")
}

func (ec *execContext) fceiling_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("fceiling")
}

func (ec *execContext) ffloor_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ffloor")
}

func (ec *execContext) fround_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("fround")
}

func (ec *execContext) ftruncate_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ftruncate")
}

func (ec *execContext) w32BatteryStatus_autogen() (lispObject, error) {
	return ec.stub("w32-battery-status")
}

func (ec *execContext) compSubrSignature_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("comp--subr-signature")
}

func (ec *execContext) compElToElnRelFilename_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("comp-el-to-eln-rel-filename")
}

func (ec *execContext) compElToElnFilename_autogen(filename, base_dir lispObject) (lispObject, error) {
	return ec.stub("comp-el-to-eln-filename")
}

func (ec *execContext) compInstallTrampoline_autogen(subr_name, trampoline lispObject) (lispObject, error) {
	return ec.stub("comp--install-trampoline")
}

func (ec *execContext) compInitCtxt_autogen() (lispObject, error) {
	return ec.stub("comp--init-ctxt")
}

func (ec *execContext) compReleaseCtxt_autogen() (lispObject, error) {
	return ec.stub("comp--release-ctxt")
}

func (ec *execContext) compNativeDriverOptionsEffectiveP_autogen() (lispObject, error) {
	return ec.stub("comp-native-driver-options-effective-p")
}

func (ec *execContext) compNativeCompilerOptionsEffectiveP_autogen() (lispObject, error) {
	return ec.stub("comp-native-compiler-options-effective-p")
}

func (ec *execContext) compCompileCtxtToFile_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("comp--compile-ctxt-to-file")
}

func (ec *execContext) compLibgccjitVersion_autogen() (lispObject, error) {
	return ec.stub("comp-libgccjit-version")
}

func (ec *execContext) compRegisterLambda_autogen(reloc_idx, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--register-lambda")
}

func (ec *execContext) compRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--register-subr")
}

func (ec *execContext) compLateRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--late-register-subr")
}

func (ec *execContext) nativeElispLoad_autogen(filename, late_load lispObject) (lispObject, error) {
	return ec.stub("native-elisp-load")
}

func (ec *execContext) nativeCompAvailableP_autogen() (lispObject, error) {
	return ec.stub("native-comp-available-p")
}

func (ec *execContext) dumpEmacsPortable_autogen(filename, track_referrers lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable")
}

func (ec *execContext) dumpEmacsPortableSortPredicate_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable--sort-predicate")
}

func (ec *execContext) dumpEmacsPortableSortPredicateCopied_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable--sort-predicate-copied")
}

func (ec *execContext) pdumperStats_autogen() (lispObject, error) {
	return ec.stub("pdumper-stats")
}

func (ec *execContext) invocationName_autogen() (lispObject, error) {
	return ec.stub("invocation-name")
}

func (ec *execContext) invocationDirectory_autogen() (lispObject, error) {
	return ec.stub("invocation-directory")
}

func (ec *execContext) killEmacs_autogen(arg, restart lispObject) (lispObject, error) {
	return ec.stub("kill-emacs")
}

func (ec *execContext) dumpEmacs_autogen(filename, symfile lispObject) (lispObject, error) {
	return ec.stub("dump-emacs")
}

func (ec *execContext) daemonp_autogen() (lispObject, error) {
	return ec.stub("daemonp")
}

func (ec *execContext) daemonInitialized_autogen() (lispObject, error) {
	return ec.stub("daemon-initialized")
}

func (ec *execContext) treesitLanguageAvailableP_autogen(language, detail lispObject) (lispObject, error) {
	return ec.stub("treesit-language-available-p")
}

func (ec *execContext) treesitLibraryAbiVersion_autogen(min_compatible lispObject) (lispObject, error) {
	return ec.stub("treesit-library-abi-version")
}

func (ec *execContext) treesitLanguageAbiVersion_autogen(language lispObject) (lispObject, error) {
	return ec.stub("treesit-language-abi-version")
}

func (ec *execContext) treesitParserP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-p")
}

func (ec *execContext) treesitNodeP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-node-p")
}

func (ec *execContext) treesitCompiledQueryP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-compiled-query-p")
}

func (ec *execContext) treesitQueryP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-query-p")
}

func (ec *execContext) treesitQueryLanguage_autogen(query lispObject) (lispObject, error) {
	return ec.stub("treesit-query-language")
}

func (ec *execContext) treesitNodeParser_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-parser")
}

func (ec *execContext) treesitParserCreate_autogen(language, buffer, no_reuse lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-create")
}

func (ec *execContext) treesitParserDelete_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-delete")
}

func (ec *execContext) treesitParserList_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-list")
}

func (ec *execContext) treesitParserBuffer_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-buffer")
}

func (ec *execContext) treesitParserLanguage_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-language")
}

func (ec *execContext) treesitParserRootNode_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-root-node")
}

func (ec *execContext) treesitParserSetIncludedRanges_autogen(parser, ranges lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-set-included-ranges")
}

func (ec *execContext) treesitParserIncludedRanges_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-included-ranges")
}

func (ec *execContext) treesitParserNotifiers_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-notifiers")
}

func (ec *execContext) treesitParserAddNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-add-notifier")
}

func (ec *execContext) treesitParserRemoveNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-remove-notifier")
}

func (ec *execContext) treesitNodeType_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-type")
}

func (ec *execContext) treesitNodeStart_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-start")
}

func (ec *execContext) treesitNodeEnd_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-end")
}

func (ec *execContext) treesitNodeString_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-string")
}

func (ec *execContext) treesitNodeParent_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-parent")
}

func (ec *execContext) treesitNodeChild_autogen(node, n, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child")
}

func (ec *execContext) treesitNodeCheck_autogen(node, property lispObject) (lispObject, error) {
	return ec.stub("treesit-node-check")
}

func (ec *execContext) treesitNodeFieldNameForChild_autogen(node, n lispObject) (lispObject, error) {
	return ec.stub("treesit-node-field-name-for-child")
}

func (ec *execContext) treesitNodeChildCount_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child-count")
}

func (ec *execContext) treesitNodeChildByFieldName_autogen(node, field_name lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child-by-field-name")
}

func (ec *execContext) treesitNodeNextSibling_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-next-sibling")
}

func (ec *execContext) treesitNodePrevSibling_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-prev-sibling")
}

func (ec *execContext) treesitNodeFirstChildForPos_autogen(node, pos, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-first-child-for-pos")
}

func (ec *execContext) treesitNodeDescendantForRange_autogen(node, beg, end, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-descendant-for-range")
}

func (ec *execContext) treesitNodeEq_autogen(node1, node2 lispObject) (lispObject, error) {
	return ec.stub("treesit-node-eq")
}

func (ec *execContext) treesitPatternExpand_autogen(pattern lispObject) (lispObject, error) {
	return ec.stub("treesit-pattern-expand")
}

func (ec *execContext) treesitQueryExpand_autogen(query lispObject) (lispObject, error) {
	return ec.stub("treesit-query-expand")
}

func (ec *execContext) treesitQueryCompile_autogen(language, query, eager lispObject) (lispObject, error) {
	return ec.stub("treesit-query-compile")
}

func (ec *execContext) treesitQueryCapture_autogen(node, query, beg, end, node_only lispObject) (lispObject, error) {
	return ec.stub("treesit-query-capture")
}

func (ec *execContext) treesitSearchSubtree_autogen(node, predicate, backward, all, depth lispObject) (lispObject, error) {
	return ec.stub("treesit-search-subtree")
}

func (ec *execContext) treesitSearchForward_autogen(start, predicate, backward, all lispObject) (lispObject, error) {
	return ec.stub("treesit-search-forward")
}

func (ec *execContext) treesitInduceSparseTree_autogen(root, predicate, process_fn, depth lispObject) (lispObject, error) {
	return ec.stub("treesit-induce-sparse-tree")
}

func (ec *execContext) treesitSubtreeStat_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-subtree-stat")
}

func (ec *execContext) treesitAvailableP_autogen() (lispObject, error) {
	return ec.stub("treesit-available-p")
}

func (ec *execContext) fontp_autogen(object, extra_type lispObject) (lispObject, error) {
	return ec.stub("fontp")
}

func (ec *execContext) fontSpec_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("font-spec")
}

func (ec *execContext) fontGet_autogen(font, key lispObject) (lispObject, error) {
	return ec.stub("font-get")
}

func (ec *execContext) fontFaceAttributes_autogen(font, frame lispObject) (lispObject, error) {
	return ec.stub("font-face-attributes")
}

func (ec *execContext) fontPut_autogen(font, prop, val lispObject) (lispObject, error) {
	return ec.stub("font-put")
}

func (ec *execContext) listFonts_autogen(font_spec, frame, num, prefer lispObject) (lispObject, error) {
	return ec.stub("list-fonts")
}

func (ec *execContext) fontFamilyList_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("font-family-list")
}

func (ec *execContext) findFont_autogen(font_spec, frame lispObject) (lispObject, error) {
	return ec.stub("find-font")
}

func (ec *execContext) fontXlfdName_autogen(font, fold_wildcards lispObject) (lispObject, error) {
	return ec.stub("font-xlfd-name")
}

func (ec *execContext) clearFontCache_autogen() (lispObject, error) {
	return ec.stub("clear-font-cache")
}

func (ec *execContext) fontShapeGstring_autogen(gstring, direction lispObject) (lispObject, error) {
	return ec.stub("font-shape-gstring")
}

func (ec *execContext) fontVariationGlyphs_autogen(font_object, character lispObject) (lispObject, error) {
	return ec.stub("font-variation-glyphs")
}

func (ec *execContext) internalCharFont_autogen(position, ch lispObject) (lispObject, error) {
	return ec.stub("internal-char-font")
}

func (ec *execContext) fontDriveOtf_autogen(otf_features, gstring_in, from, to, gstring_out, index lispObject) (lispObject, error) {
	return ec.stub("font-drive-otf")
}

func (ec *execContext) fontOtfAlternates_autogen(font_object, character, otf_features lispObject) (lispObject, error) {
	return ec.stub("font-otf-alternates")
}

func (ec *execContext) openFont_autogen(font_entity, size, frame lispObject) (lispObject, error) {
	return ec.stub("open-font")
}

func (ec *execContext) closeFont_autogen(font_object, frame lispObject) (lispObject, error) {
	return ec.stub("close-font")
}

func (ec *execContext) queryFont_autogen(font_object lispObject) (lispObject, error) {
	return ec.stub("query-font")
}

func (ec *execContext) fontHasCharP_autogen(font, ch, frame lispObject) (lispObject, error) {
	return ec.stub("font-has-char-p")
}

func (ec *execContext) fontGetGlyphs_autogen(font_object, from, to, object lispObject) (lispObject, error) {
	return ec.stub("font-get-glyphs")
}

func (ec *execContext) fontMatchP_autogen(spec, font lispObject) (lispObject, error) {
	return ec.stub("font-match-p")
}

func (ec *execContext) fontAt_autogen(position, window, string lispObject) (lispObject, error) {
	return ec.stub("font-at")
}

func (ec *execContext) drawString_autogen(font_object, string lispObject) (lispObject, error) {
	return ec.stub("draw-string")
}

func (ec *execContext) frameFontCache_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-font-cache")
}

func (ec *execContext) fontInfo_autogen(name, frame lispObject) (lispObject, error) {
	return ec.stub("font-info")
}

func (ec *execContext) moduleLoad_autogen(file lispObject) (lispObject, error) {
	return ec.stub("module-load")
}

func (ec *execContext) playSoundInternal_autogen(sound lispObject) (lispObject, error) {
	return ec.stub("play-sound-internal")
}

func (ec *execContext) imageSize_autogen(spec, pixels, frame lispObject) (lispObject, error) {
	return ec.stub("image-size")
}

func (ec *execContext) imageMaskP_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-mask-p")
}

func (ec *execContext) imageMetadata_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-metadata")
}

func (ec *execContext) clearImageCache_autogen(filter, animation_cache lispObject) (lispObject, error) {
	return ec.stub("clear-image-cache")
}

func (ec *execContext) imageFlush_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-flush")
}

func (ec *execContext) imagemagickTypes_autogen() (lispObject, error) {
	return ec.stub("imagemagick-types")
}

func (ec *execContext) imagep_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("imagep")
}

func (ec *execContext) lookupImage_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("lookup-image")
}

func (ec *execContext) imageTransformsP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("image-transforms-p")
}

func (ec *execContext) imageCacheSize_autogen() (lispObject, error) {
	return ec.stub("image-cache-size")
}

func (ec *execContext) initImageLibrary_autogen(type_ lispObject) (lispObject, error) {
	return ec.stub("init-image-library")
}

func (ec *execContext) gfileAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return ec.stub("gfile-add-watch")
}

func (ec *execContext) gfileRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-rm-watch")
}

func (ec *execContext) gfileValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-valid-p")
}

func (ec *execContext) gfileMonitorName_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-monitor-name")
}

func (ec *execContext) dbusInitBus_autogen(bus, private lispObject) (lispObject, error) {
	return ec.stub("dbus--init-bus")
}

func (ec *execContext) dbusGetUniqueName_autogen(bus lispObject) (lispObject, error) {
	return ec.stub("dbus-get-unique-name")
}

func (ec *execContext) dbusMessageInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("dbus-message-internal")
}

func (ec *execContext) destroyFringeBitmap_autogen(bitmap lispObject) (lispObject, error) {
	return ec.stub("destroy-fringe-bitmap")
}

func (ec *execContext) defineFringeBitmap_autogen(bitmap, bits, height, width, align lispObject) (lispObject, error) {
	return ec.stub("define-fringe-bitmap")
}

func (ec *execContext) setFringeBitmapFace_autogen(bitmap, face lispObject) (lispObject, error) {
	return ec.stub("set-fringe-bitmap-face")
}

func (ec *execContext) fringeBitmapsAtPos_autogen(pos, window lispObject) (lispObject, error) {
	return ec.stub("fringe-bitmaps-at-pos")
}

func (ec *execContext) framep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("framep")
}

func (ec *execContext) frameLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("frame-live-p")
}

func (ec *execContext) windowSystem_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("window-system")
}

func (ec *execContext) frameWindowsMinSize_autogen(frame, horizontal, ignore, pixelwise lispObject) (lispObject, error) {
	return ec.stub("frame-windows-min-size")
}

func (ec *execContext) makeTerminalFrame_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("make-terminal-frame")
}

func (ec *execContext) selectFrame_autogen(frame, norecord lispObject) (lispObject, error) {
	return ec.stub("select-frame")
}

func (ec *execContext) handleSwitchFrame_autogen(event lispObject) (lispObject, error) {
	return ec.stub("handle-switch-frame")
}

func (ec *execContext) selectedFrame_autogen() (lispObject, error) {
	return ec.stub("selected-frame")
}

func (ec *execContext) oldSelectedFrame_autogen() (lispObject, error) {
	return ec.stub("old-selected-frame")
}

func (ec *execContext) frameList_autogen() (lispObject, error) {
	return ec.stub("frame-list")
}

func (ec *execContext) frameParent_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-parent")
}

func (ec *execContext) frameAncestorP_autogen(ancestor, descendant lispObject) (lispObject, error) {
	return ec.stub("frame-ancestor-p")
}

func (ec *execContext) nextFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return ec.stub("next-frame")
}

func (ec *execContext) previousFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return ec.stub("previous-frame")
}

func (ec *execContext) lastNonminibufFrame_autogen() (lispObject, error) {
	return ec.stub("last-nonminibuffer-frame")
}

func (ec *execContext) deleteFrame_autogen(frame, force lispObject) (lispObject, error) {
	return ec.stub("delete-frame")
}

func (ec *execContext) mousePosition_autogen() (lispObject, error) {
	return ec.stub("mouse-position")
}

func (ec *execContext) mousePixelPosition_autogen() (lispObject, error) {
	return ec.stub("mouse-pixel-position")
}

func (ec *execContext) setMousePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-mouse-position")
}

func (ec *execContext) setMousePixelPosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-mouse-pixel-position")
}

func (ec *execContext) makeFrameVisible_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("make-frame-visible")
}

func (ec *execContext) makeFrameInvisible_autogen(frame, force lispObject) (lispObject, error) {
	return ec.stub("make-frame-invisible")
}

func (ec *execContext) iconifyFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("iconify-frame")
}

func (ec *execContext) frameVisibleP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-visible-p")
}

func (ec *execContext) visibleFrameList_autogen() (lispObject, error) {
	return ec.stub("visible-frame-list")
}

func (ec *execContext) raiseFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("raise-frame")
}

func (ec *execContext) lowerFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("lower-frame")
}

func (ec *execContext) redirectFrameFocus_autogen(frame, focus_frame lispObject) (lispObject, error) {
	return ec.stub("redirect-frame-focus")
}

func (ec *execContext) frameFocus_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-focus")
}

func (ec *execContext) xFocusFrame_autogen(frame, noactivate lispObject) (lispObject, error) {
	return ec.stub("x-focus-frame")
}

func (ec *execContext) frameAfterMakeFrame_autogen(frame, made lispObject) (lispObject, error) {
	return ec.stub("frame-after-make-frame")
}

func (ec *execContext) frameParameters_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-parameters")
}

func (ec *execContext) frameParameter_autogen(frame, parameter lispObject) (lispObject, error) {
	return ec.stub("frame-parameter")
}

func (ec *execContext) modifyFrameParameters_autogen(frame, alist lispObject) (lispObject, error) {
	return ec.stub("modify-frame-parameters")
}

func (ec *execContext) frameCharHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-char-height")
}

func (ec *execContext) frameCharWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-char-width")
}

func (ec *execContext) frameNativeWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-native-width")
}

func (ec *execContext) frameNativeHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-native-height")
}

func (ec *execContext) toolBarPixelWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("tool-bar-pixel-width")
}

func (ec *execContext) frameTextCols_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-cols")
}

func (ec *execContext) frameTextLines_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-lines")
}

func (ec *execContext) frameTotalCols_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-total-cols")
}

func (ec *execContext) frameTotalLines_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-total-lines")
}

func (ec *execContext) frameTextWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-width")
}

func (ec *execContext) frameTextHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-height")
}

func (ec *execContext) scrollBarWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scroll-bar-width")
}

func (ec *execContext) scrollBarHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scroll-bar-height")
}

func (ec *execContext) fringeWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-fringe-width")
}

func (ec *execContext) frameChildFrameBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-child-frame-border-width")
}

func (ec *execContext) frameInternalBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-internal-border-width")
}

func (ec *execContext) rightDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-right-divider-width")
}

func (ec *execContext) bottomDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-bottom-divider-width")
}

func (ec *execContext) setFrameHeight_autogen(frame, height, pretend, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-height")
}

func (ec *execContext) setFrameWidth_autogen(frame, width, pretend, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-width")
}

func (ec *execContext) setFrameSize_autogen(frame, width, height, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-size")
}

func (ec *execContext) framePosition_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-position")
}

func (ec *execContext) setFramePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-frame-position")
}

func (ec *execContext) frameWindowStateChange_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-window-state-change")
}

func (ec *execContext) setFrameWindowStateChange_autogen(frame, arg lispObject) (lispObject, error) {
	return ec.stub("set-frame-window-state-change")
}

func (ec *execContext) frameScaleFactor_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scale-factor")
}

func (ec *execContext) xGetResource_autogen(attribute, class, component, subclass lispObject) (lispObject, error) {
	return ec.stub("x-get-resource")
}

func (ec *execContext) xParseGeometry_autogen(string lispObject) (lispObject, error) {
	return ec.stub("x-parse-geometry")
}

func (ec *execContext) framePointerVisibleP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-pointer-visible-p")
}

func (ec *execContext) frameSetWasInvisible_autogen(frame, was_invisible lispObject) (lispObject, error) {
	return ec.stub("frame--set-was-invisible")
}

func (ec *execContext) reconsiderFrameFonts_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("reconsider-frame-fonts")
}

func (ec *execContext) eq_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return ec.stub("eq")
}

func (ec *execContext) null_autogen(object lispObject) (lispObject, error) {
	return ec.stub("null")
}

func (ec *execContext) typeOf_autogen(object lispObject) (lispObject, error) {
	return ec.stub("type-of")
}

func (ec *execContext) consp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("consp")
}

func (ec *execContext) atom_autogen(object lispObject) (lispObject, error) {
	return ec.stub("atom")
}

func (ec *execContext) listp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("listp")
}

func (ec *execContext) nlistp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("nlistp")
}

func (ec *execContext) bareSymbolP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bare-symbol-p")
}

func (ec *execContext) symbolWithPosP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("symbol-with-pos-p")
}

func (ec *execContext) symbolp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("symbolp")
}

func (ec *execContext) keywordp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("keywordp")
}

func (ec *execContext) vectorp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("vectorp")
}

func (ec *execContext) recordp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("recordp")
}

func (ec *execContext) stringp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("stringp")
}

func (ec *execContext) multibyteStringP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("multibyte-string-p")
}

func (ec *execContext) charTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("char-table-p")
}

func (ec *execContext) vectorOrCharTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("vector-or-char-table-p")
}

func (ec *execContext) boolVectorP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bool-vector-p")
}

func (ec *execContext) arrayp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("arrayp")
}

func (ec *execContext) sequencep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("sequencep")
}

func (ec *execContext) bufferp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bufferp")
}

func (ec *execContext) markerp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("markerp")
}

func (ec *execContext) userPtrp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("user-ptrp")
}

func (ec *execContext) subrp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("subrp")
}

func (ec *execContext) byteCodeFunctionP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("byte-code-function-p")
}

func (ec *execContext) moduleFunctionP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("module-function-p")
}

func (ec *execContext) charOrStringP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("char-or-string-p")
}

func (ec *execContext) integerp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("integerp")
}

func (ec *execContext) integerOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("integer-or-marker-p")
}

func (ec *execContext) natnump_autogen(object lispObject) (lispObject, error) {
	return ec.stub("natnump")
}

func (ec *execContext) numberp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("numberp")
}

func (ec *execContext) numberOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("number-or-marker-p")
}

func (ec *execContext) floatp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("floatp")
}

func (ec *execContext) threadp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("threadp")
}

func (ec *execContext) mutexp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("mutexp")
}

func (ec *execContext) conditionVariableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("condition-variable-p")
}

func (ec *execContext) car_autogen(list lispObject) (lispObject, error) {
	return ec.stub("car")
}

func (ec *execContext) carSafe_autogen(object lispObject) (lispObject, error) {
	return ec.stub("car-safe")
}

func (ec *execContext) cdr_autogen(list lispObject) (lispObject, error) {
	return ec.stub("cdr")
}

func (ec *execContext) cdrSafe_autogen(object lispObject) (lispObject, error) {
	return ec.stub("cdr-safe")
}

func (ec *execContext) setcar_autogen(cell, newcar lispObject) (lispObject, error) {
	return ec.stub("setcar")
}

func (ec *execContext) setcdr_autogen(cell, newcdr lispObject) (lispObject, error) {
	return ec.stub("setcdr")
}

func (ec *execContext) boundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("boundp")
}

func (ec *execContext) fboundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("fboundp")
}

func (ec *execContext) makunbound_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("makunbound")
}

func (ec *execContext) fmakunbound_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("fmakunbound")
}

func (ec *execContext) symbolFunction_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-function")
}

func (ec *execContext) symbolPlist_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-plist")
}

func (ec *execContext) symbolName_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-name")
}

func (ec *execContext) bareSymbol_autogen(sym lispObject) (lispObject, error) {
	return ec.stub("bare-symbol")
}

func (ec *execContext) symbolWithPosPos_autogen(ls lispObject) (lispObject, error) {
	return ec.stub("symbol-with-pos-pos")
}

func (ec *execContext) removePosFromSymbol_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("remove-pos-from-symbol")
}

func (ec *execContext) positionSymbol_autogen(sym, pos lispObject) (lispObject, error) {
	return ec.stub("position-symbol")
}

func (ec *execContext) fset_autogen(symbol, definition lispObject) (lispObject, error) {
	return ec.stub("fset")
}

func (ec *execContext) defalias_autogen(symbol, definition, docstring lispObject) (lispObject, error) {
	return ec.stub("defalias")
}

func (ec *execContext) setplist_autogen(symbol, newplist lispObject) (lispObject, error) {
	return ec.stub("setplist")
}

func (ec *execContext) subrArity_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-arity")
}

func (ec *execContext) subrName_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-name")
}

func (ec *execContext) subrNativeElispP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("subr-native-elisp-p")
}

func (ec *execContext) subrNativeLambdaList_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-native-lambda-list")
}

func (ec *execContext) subrType_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-type")
}

func (ec *execContext) subrNativeCompUnit_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-native-comp-unit")
}

func (ec *execContext) nativeCompUnitFile_autogen(comp_unit lispObject) (lispObject, error) {
	return ec.stub("native-comp-unit-file")
}

func (ec *execContext) nativeCompUnitSetFile_autogen(comp_unit, new_file lispObject) (lispObject, error) {
	return ec.stub("native-comp-unit-set-file")
}

func (ec *execContext) interactiveForm_autogen(cmd lispObject) (lispObject, error) {
	return ec.stub("interactive-form")
}

func (ec *execContext) commandModes_autogen(command lispObject) (lispObject, error) {
	return ec.stub("command-modes")
}

func (ec *execContext) indirectVariable_autogen(object lispObject) (lispObject, error) {
	return ec.stub("indirect-variable")
}

func (ec *execContext) symbolValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-value")
}

func (ec *execContext) set_autogen(symbol, newval lispObject) (lispObject, error) {
	return ec.stub("set")
}

func (ec *execContext) addVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return ec.stub("add-variable-watcher")
}

func (ec *execContext) removeVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return ec.stub("remove-variable-watcher")
}

func (ec *execContext) getVariableWatchers_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("get-variable-watchers")
}

func (ec *execContext) defaultBoundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-boundp")
}

func (ec *execContext) defaultValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-value")
}

func (ec *execContext) setDefault_autogen(symbol, value lispObject) (lispObject, error) {
	return ec.stub("set-default")
}

func (ec *execContext) makeVariableBufferLocal_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("make-variable-buffer-local")
}

func (ec *execContext) makeLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("make-local-variable")
}

func (ec *execContext) killLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("kill-local-variable")
}

func (ec *execContext) localVariableP_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("local-variable-p")
}

func (ec *execContext) localVariableIfSetP_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("local-variable-if-set-p")
}

func (ec *execContext) variableBindingLocus_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("variable-binding-locus")
}

func (ec *execContext) indirectFunction_autogen(object, noerror lispObject) (lispObject, error) {
	return ec.stub("indirect-function")
}

func (ec *execContext) aref_autogen(array, idx lispObject) (lispObject, error) {
	return ec.stub("aref")
}

func (ec *execContext) aset_autogen(array, idx, newelt lispObject) (lispObject, error) {
	return ec.stub("aset")
}

func (ec *execContext) eqlsign_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("=")
}

func (ec *execContext) lss_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("<")
}

func (ec *execContext) gtr_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub(">")
}

func (ec *execContext) leq_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("<=")
}

func (ec *execContext) geq_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub(">=")
}

func (ec *execContext) neq_autogen(num1, num2 lispObject) (lispObject, error) {
	return ec.stub("/=")
}

func (ec *execContext) numberToString_autogen(number lispObject) (lispObject, error) {
	return ec.stub("number-to-string")
}

func (ec *execContext) stringToNumber_autogen(string, base lispObject) (lispObject, error) {
	return ec.stub("string-to-number")
}

func (ec *execContext) plus_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("+")
}

func (ec *execContext) minus_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("-")
}

func (ec *execContext) times_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("*")
}

func (ec *execContext) quo_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("/")
}

func (ec *execContext) rem_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("%")
}

func (ec *execContext) mod_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("mod")
}

func (ec *execContext) max_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("max")
}

func (ec *execContext) min_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("min")
}

func (ec *execContext) logand_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logand")
}

func (ec *execContext) logior_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logior")
}

func (ec *execContext) logxor_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logxor")
}

func (ec *execContext) logcount_autogen(value lispObject) (lispObject, error) {
	return ec.stub("logcount")
}

func (ec *execContext) ash_autogen(value, count lispObject) (lispObject, error) {
	return ec.stub("ash")
}

func (ec *execContext) add1_autogen(number lispObject) (lispObject, error) {
	return ec.stub("1+")
}

func (ec *execContext) sub1_autogen(number lispObject) (lispObject, error) {
	return ec.stub("1-")
}

func (ec *execContext) lognot_autogen(number lispObject) (lispObject, error) {
	return ec.stub("lognot")
}

func (ec *execContext) byteorder_autogen() (lispObject, error) {
	return ec.stub("byteorder")
}

func (ec *execContext) boolVectorExclusiveOr_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-exclusive-or")
}

func (ec *execContext) boolVectorUnion_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-union")
}

func (ec *execContext) boolVectorIntersection_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-intersection")
}

func (ec *execContext) boolVectorSetDifference_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-set-difference")
}

func (ec *execContext) boolVectorSubsetp_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("bool-vector-subsetp")
}

func (ec *execContext) boolVectorNot_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("bool-vector-not")
}

func (ec *execContext) boolVectorCountPopulation_autogen(a lispObject) (lispObject, error) {
	return ec.stub("bool-vector-count-population")
}

func (ec *execContext) boolVectorCountConsecutive_autogen(a, b, i lispObject) (lispObject, error) {
	return ec.stub("bool-vector-count-consecutive")
}

func (ec *execContext) profilerCpuStart_autogen(sampling_interval lispObject) (lispObject, error) {
	return ec.stub("profiler-cpu-start")
}

func (ec *execContext) profilerCpuStop_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-stop")
}

func (ec *execContext) profilerCpuRunningP_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-running-p")
}

func (ec *execContext) profilerCpuLog_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-log")
}

func (ec *execContext) profilerMemoryStart_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-start")
}

func (ec *execContext) profilerMemoryStop_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-stop")
}

func (ec *execContext) profilerMemoryRunningP_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-running-p")
}

func (ec *execContext) profilerMemoryLog_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-log")
}

func (ec *execContext) functionEqual_autogen(f1, f2 lispObject) (lispObject, error) {
	return ec.stub("function-equal")
}

func (ec *execContext) lockFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("lock-file")
}

func (ec *execContext) unlockFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("unlock-file")
}

func (ec *execContext) lockBuffer_autogen(file lispObject) (lispObject, error) {
	return ec.stub("lock-buffer")
}

func (ec *execContext) unlockBuffer_autogen() (lispObject, error) {
	return ec.stub("unlock-buffer")
}

func (ec *execContext) fileLockedP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-locked-p")
}

func (ec *execContext) symbolsOfEmacs_autogen() {
	ec.defSubr0("active-minibuffer-window", ec.activeMinibufferWindow_autogen)
	ec.defSubr1("set-minibuffer-window", ec.setMinibufferWindow_autogen, 1)
	ec.defSubr2("minibufferp", ec.minibufferp_autogen, 0)
	ec.defSubr1("innermost-minibuffer-p", ec.innermostMinibufferP_autogen, 0)
	ec.defSubr1("minibuffer-innermost-command-loop-p", ec.minibufferInnermostCommandLoopP_autogen, 0)
	ec.defSubr0("abort-minibuffers", ec.abortMinibuffers_autogen)
	ec.defSubr0("minibuffer-prompt-end", ec.minibufferPromptEnd_autogen)
	ec.defSubr0("minibuffer-contents", ec.minibufferContents_autogen)
	ec.defSubr0("minibuffer-contents-no-properties", ec.minibufferContentsNoProperties_autogen)
	ec.defSubr7("read-from-minibuffer", ec.readFromMinibuffer_autogen, 1)
	ec.defSubr5("read-string", ec.readString_autogen, 1)
	ec.defSubr2("read-command", ec.readCommand_autogen, 1)
	ec.defSubr1("read-function", ec.readFunction_autogen, 1)
	ec.defSubr2("read-variable", ec.readVariable_autogen, 1)
	ec.defSubr4("read-buffer", ec.readBuffer_autogen, 1)
	ec.defSubr3("try-completion", ec.tryCompletion_autogen, 2)
	ec.defSubr4("all-completions", ec.allCompletions_autogen, 2)
	ec.defSubr8("completing-read", ec.completingRead_autogen, 2)
	ec.defSubr3("test-completion", ec.testCompletion_autogen, 2)
	ec.defSubr3("internal-complete-buffer", ec.internalCompleteBuffer_autogen, 3)
	ec.defSubr3("assoc-string", ec.assocString_autogen, 2)
	ec.defSubr0("minibuffer-depth", ec.minibufferDepth_autogen)
	ec.defSubr0("minibuffer-prompt", ec.minibufferPrompt_autogen)
	ec.defSubr2("time-add", ec.timeAdd_autogen, 2)
	ec.defSubr2("time-subtract", ec.timeSubtract_autogen, 2)
	ec.defSubr2("time-less-p", ec.timeLessP_autogen, 2)
	ec.defSubr2("time-equal-p", ec.timeEqualP_autogen, 2)
	ec.defSubr1("float-time", ec.floatTime_autogen, 0)
	ec.defSubr3("format-time-string", ec.formatTimeString_autogen, 1)
	ec.defSubr3("decode-time", ec.decodeTime_autogen, 0)
	ec.defSubrM("encode-time", ec.encodeTime_autogen, 1)
	ec.defSubr2("time-convert", ec.timeConvert_autogen, 1)
	ec.defSubr0("current-time", ec.currentTime_autogen)
	ec.defSubr0("current-cpu-time", ec.currentCpuTime_autogen)
	ec.defSubr2("current-time-string", ec.currentTimeString_autogen, 0)
	ec.defSubr2("current-time-zone", ec.currentTimeZone_autogen, 0)
	ec.defSubr1("set-time-zone-rule", ec.setTimeZoneRule_autogen, 1)
	ec.defSubr1("tty-display-color-p", ec.ttyDisplayColorP_autogen, 0)
	ec.defSubr1("tty-display-color-cells", ec.ttyDisplayColorCells_autogen, 0)
	ec.defSubr1("tty-type", ec.ttyType_autogen, 0)
	ec.defSubr1("controlling-tty-p", ec.controllingTtyP_autogen, 0)
	ec.defSubr1("tty-no-underline", ec.ttyNoUnderline_autogen, 0)
	ec.defSubr1("tty-top-frame", ec.ttyTopFrame_autogen, 0)
	ec.defSubr1("suspend-tty", ec.suspendTty_autogen, 0)
	ec.defSubr1("resume-tty", ec.resumeTty_autogen, 0)
	ec.defSubr2("tty--set-output-buffer-size", ec.ttySetOutputBufferSize_autogen, 1)
	ec.defSubr1("tty--output-buffer-size", ec.ttyOutputBufferSize_autogen, 0)
	ec.defSubr0("gpm-mouse-start", ec.gpmMouseStart_autogen)
	ec.defSubr0("gpm-mouse-stop", ec.gpmMouseStop_autogen)
	ec.defSubr0("undo-boundary", ec.undoBoundary_autogen)
	ec.defSubr2("start-kbd-macro", ec.startKbdMacro_autogen, 1)
	ec.defSubr2("end-kbd-macro", ec.endKbdMacro_autogen, 0)
	ec.defSubr0("cancel-kbd-macro-events", ec.cancelKbdMacroEvents_autogen)
	ec.defSubr1("store-kbd-macro-event", ec.storeKbdMacroEvent_autogen, 1)
	ec.defSubr2("call-last-kbd-macro", ec.callLastKbdMacro_autogen, 0)
	ec.defSubr3("execute-kbd-macro", ec.executeKbdMacro_autogen, 1)
	ec.defSubr4("libxml-parse-html-region", ec.libxmlParseHtmlRegion_autogen, 0)
	ec.defSubr4("libxml-parse-xml-region", ec.libxmlParseXmlRegion_autogen, 0)
	ec.defSubr0("libxml-available-p", ec.libxmlAvailableP_autogen)
	ec.defSubr3("kqueue-add-watch", ec.kqueueAddWatch_autogen, 3)
	ec.defSubr1("kqueue-rm-watch", ec.kqueueRmWatch_autogen, 1)
	ec.defSubr1("kqueue-valid-p", ec.kqueueValidP_autogen, 1)
	ec.defSubr0("combine-after-change-execute", ec.combineAfterChangeExecute_autogen)
	ec.defSubr1("ccl-program-p", ec.cclProgramP_autogen, 1)
	ec.defSubr2("ccl-execute", ec.cclExecute_autogen, 2)
	ec.defSubr5("ccl-execute-on-string", ec.cclExecuteOnString_autogen, 3)
	ec.defSubr2("register-ccl-program", ec.registerCclProgram_autogen, 2)
	ec.defSubr2("register-code-conversion-map", ec.registerCodeConversionMap_autogen, 2)
	ec.defSubr1("syntax-table-p", ec.syntaxTableP_autogen, 1)
	ec.defSubr0("syntax-table", ec.syntaxTable_autogen)
	ec.defSubr0("standard-syntax-table", ec.standardSyntaxTable_autogen)
	ec.defSubr1("copy-syntax-table", ec.copySyntaxTable_autogen, 0)
	ec.defSubr1("set-syntax-table", ec.setSyntaxTable_autogen, 1)
	ec.defSubr1("char-syntax", ec.charSyntax_autogen, 1)
	ec.defSubr1("syntax-class-to-char", ec.syntaxClassToChar_autogen, 1)
	ec.defSubr1("matching-paren", ec.matchingParen_autogen, 1)
	ec.defSubr1("string-to-syntax", ec.stringToSyntax_autogen, 1)
	ec.defSubr3("modify-syntax-entry", ec.modifySyntaxEntry_autogen, 2)
	ec.defSubr1("internal-describe-syntax-value", ec.internalDescribeSyntaxValue_autogen, 1)
	ec.defSubr1("forward-word", ec.forwardWord_autogen, 0)
	ec.defSubr2("skip-chars-forward", ec.skipCharsForward_autogen, 1)
	ec.defSubr2("skip-chars-backward", ec.skipCharsBackward_autogen, 1)
	ec.defSubr2("skip-syntax-forward", ec.skipSyntaxForward_autogen, 1)
	ec.defSubr2("skip-syntax-backward", ec.skipSyntaxBackward_autogen, 1)
	ec.defSubr1("forward-comment", ec.forwardComment_autogen, 1)
	ec.defSubr3("scan-lists", ec.scanLists_autogen, 3)
	ec.defSubr2("scan-sexps", ec.scanSexps_autogen, 2)
	ec.defSubr0("backward-prefix-chars", ec.backwardPrefixChars_autogen)
	ec.defSubr6("parse-partial-sexp", ec.parsePartialSexp_autogen, 2)
	ec.defSubr0("clear-composition-cache", ec.clearCompositionCache_autogen)
	ec.defSubr4("composition-get-gstring", ec.compositionGetGstring_autogen, 4)
	ec.defSubr4("compose-region-internal", ec.composeRegionInternal_autogen, 2)
	ec.defSubr5("compose-string-internal", ec.composeStringInternal_autogen, 3)
	ec.defSubr4("find-composition-internal", ec.findCompositionInternal_autogen, 4)
	ec.defSubr1("composition-sort-rules", ec.compositionSortRules_autogen, 1)
	ec.defSubr3("inotify-add-watch", ec.inotifyAddWatch_autogen, 3)
	ec.defSubr1("inotify-rm-watch", ec.inotifyRmWatch_autogen, 1)
	ec.defSubr1("inotify-valid-p", ec.inotifyValidP_autogen, 1)
	ec.defSubr0("inotify-watch-list", ec.inotifyWatchList_autogen)
	ec.defSubr0("inotify-allocated-p", ec.inotifyAllocatedP_autogen)
	ec.defSubr1("char-to-string", ec.charToString_autogen, 1)
	ec.defSubr1("byte-to-string", ec.byteToString_autogen, 1)
	ec.defSubr1("string-to-char", ec.stringToChar_autogen, 1)
	ec.defSubr0("point", ec.point_autogen)
	ec.defSubr0("point-marker", ec.pointMarker_autogen)
	ec.defSubr1("goto-char", ec.gotoChar_autogen, 1)
	ec.defSubr0("region-beginning", ec.regionBeginning_autogen)
	ec.defSubr0("region-end", ec.regionEnd_autogen)
	ec.defSubr0("mark-marker", ec.markMarker_autogen)
	ec.defSubr3("get-pos-property", ec.getPosProperty_autogen, 2)
	ec.defSubr1("delete-field", ec.deleteField_autogen, 0)
	ec.defSubr1("field-string", ec.fieldString_autogen, 0)
	ec.defSubr1("field-string-no-properties", ec.fieldStringNoProperties_autogen, 0)
	ec.defSubr3("field-beginning", ec.fieldBeginning_autogen, 0)
	ec.defSubr3("field-end", ec.fieldEnd_autogen, 0)
	ec.defSubr5("constrain-to-field", ec.constrainToField_autogen, 2)
	ec.defSubr1("pos-bol", ec.posBol_autogen, 0)
	ec.defSubr1("line-beginning-position", ec.lineBeginningPosition_autogen, 0)
	ec.defSubr1("pos-eol", ec.posEol_autogen, 0)
	ec.defSubr1("line-end-position", ec.lineEndPosition_autogen, 0)
	ec.defSubrU("save-excursion", ec.saveExcursion_autogen, 0)
	ec.defSubrU("save-current-buffer", ec.saveCurrentBuffer_autogen, 0)
	ec.defSubr1("buffer-size", ec.bufferSize_autogen, 0)
	ec.defSubr0("point-min", ec.pointMin_autogen)
	ec.defSubr0("point-min-marker", ec.pointMinMarker_autogen)
	ec.defSubr0("point-max", ec.pointMax_autogen)
	ec.defSubr0("point-max-marker", ec.pointMaxMarker_autogen)
	ec.defSubr0("gap-position", ec.gapPosition_autogen)
	ec.defSubr0("gap-size", ec.gapSize_autogen)
	ec.defSubr1("position-bytes", ec.positionBytes_autogen, 1)
	ec.defSubr1("byte-to-position", ec.byteToPosition_autogen, 1)
	ec.defSubr0("following-char", ec.followingChar_autogen)
	ec.defSubr0("preceding-char", ec.previousChar_autogen)
	ec.defSubr0("bobp", ec.bobp_autogen)
	ec.defSubr0("eobp", ec.eobp_autogen)
	ec.defSubr0("bolp", ec.bolp_autogen)
	ec.defSubr0("eolp", ec.eolp_autogen)
	ec.defSubr1("char-after", ec.charAfter_autogen, 0)
	ec.defSubr1("char-before", ec.charBefore_autogen, 0)
	ec.defSubr1("user-login-name", ec.userLoginName_autogen, 0)
	ec.defSubr0("user-real-login-name", ec.userRealLoginName_autogen)
	ec.defSubr0("user-uid", ec.userUid_autogen)
	ec.defSubr0("user-real-uid", ec.userRealUid_autogen)
	ec.defSubr1("group-name", ec.groupName_autogen, 1)
	ec.defSubr0("group-gid", ec.groupGid_autogen)
	ec.defSubr0("group-real-gid", ec.groupRealGid_autogen)
	ec.defSubr1("user-full-name", ec.userFullName_autogen, 0)
	ec.defSubr0("system-name", ec.systemName_autogen)
	ec.defSubr0("emacs-pid", ec.emacsPid_autogen)
	ec.defSubrM("insert", ec.insert_autogen, 0)
	ec.defSubrM("insert-and-inherit", ec.insertAndInherit_autogen, 0)
	ec.defSubrM("insert-before-markers", ec.insertBeforeMarkers_autogen, 0)
	ec.defSubrM("insert-before-markers-and-inherit", ec.insertAndInheritBeforeMarkers_autogen, 0)
	ec.defSubr3("insert-char", ec.insertChar_autogen, 1)
	ec.defSubr3("insert-byte", ec.insertByte_autogen, 2)
	ec.defSubr2("buffer-substring", ec.bufferSubstring_autogen, 2)
	ec.defSubr2("buffer-substring-no-properties", ec.bufferSubstringNoProperties_autogen, 2)
	ec.defSubr0("buffer-string", ec.bufferString_autogen)
	ec.defSubr3("insert-buffer-substring", ec.insertBufferSubstring_autogen, 1)
	ec.defSubr6("compare-buffer-substrings", ec.compareBufferSubstrings_autogen, 6)
	ec.defSubr3("replace-buffer-contents", ec.replaceBufferContents_autogen, 1)
	ec.defSubr5("subst-char-in-region", ec.substCharInRegion_autogen, 4)
	ec.defSubr3("translate-region-internal", ec.translateRegionInternal_autogen, 3)
	ec.defSubr2("delete-region", ec.deleteRegion_autogen, 2)
	ec.defSubr2("delete-and-extract-region", ec.deleteAndExtractRegion_autogen, 2)
	ec.defSubr0("widen", ec.widen_autogen)
	ec.defSubr2("narrow-to-region", ec.narrowToRegion_autogen, 2)
	ec.defSubr1("internal--lock-narrowing", ec.internalLockNarrowing_autogen, 1)
	ec.defSubr1("internal--unlock-narrowing", ec.internalUnlockNarrowing_autogen, 1)
	ec.defSubrU("save-restriction", ec.saveRestriction_autogen, 0)
	ec.defSubr3("ngettext", ec.ngettext_autogen, 3)
	ec.defSubrM("message", ec.message_autogen, 1)
	ec.defSubrM("message-box", ec.messageBox_autogen, 1)
	ec.defSubrM("message-or-box", ec.messageOrBox_autogen, 1)
	ec.defSubr0("current-message", ec.currentMessage_autogen)
	ec.defSubrM("propertize", ec.propertize_autogen, 1)
	ec.defSubrM("format", ec.format_autogen, 1)
	ec.defSubrM("format-message", ec.formatMessage_autogen, 1)
	ec.defSubr2("char-equal", ec.charEqual_autogen, 2)
	ec.defSubr5("transpose-regions", ec.transposeRegions_autogen, 4)
	ec.defSubr2("cygwin-convert-file-name-to-windows", ec.cygwinConvertFileNameToWindows_autogen, 1)
	ec.defSubr2("cygwin-convert-file-name-from-windows", ec.cygwinConvertFileNameFromWindows_autogen, 1)
	ec.defSubr0("recursive-edit", ec.recursiveEdit_autogen)
	ec.defSubr3("command-error-default-function", ec.commandErrorDefaultFunction_autogen, 3)
	ec.defSubr0("top-level", ec.topLevel_autogen)
	ec.defSubr0("exit-recursive-edit", ec.exitRecursiveEdit_autogen)
	ec.defSubr0("abort-recursive-edit", ec.abortRecursiveEdit_autogen)
	ec.defSubr1("internal--track-mouse", ec.internalTrackMouse_autogen, 1)
	ec.defSubr0("current-idle-time", ec.currentIdleTime_autogen)
	ec.defSubr1("internal-event-symbol-parse-modifiers", ec.eventSymbolParseModifiers_autogen, 1)
	ec.defSubr1("event-convert-list", ec.eventConvertList_autogen, 1)
	ec.defSubr1("internal-handle-focus-in", ec.internalHandleFocusIn_autogen, 1)
	ec.defSubr5("read-key-sequence", ec.readKeySequence_autogen, 1)
	ec.defSubr5("read-key-sequence-vector", ec.readKeySequenceVector_autogen, 1)
	ec.defSubr1("input-pending-p", ec.inputPendingP_autogen, 0)
	ec.defSubr1("lossage-size", ec.lossageSize_autogen, 0)
	ec.defSubr1("recent-keys", ec.recentKeys_autogen, 0)
	ec.defSubr0("this-command-keys", ec.thisCommandKeys_autogen)
	ec.defSubr1("set--this-command-keys", ec.setThisCommandKeys_autogen, 1)
	ec.defSubr0("this-command-keys-vector", ec.thisCommandKeysVector_autogen)
	ec.defSubr0("this-single-command-keys", ec.thisSingleCommandKeys_autogen)
	ec.defSubr0("this-single-command-raw-keys", ec.thisSingleCommandRawKeys_autogen)
	ec.defSubr1("clear-this-command-keys", ec.clearThisCommandKeys_autogen, 0)
	ec.defSubr0("recursion-depth", ec.recursionDepth_autogen)
	ec.defSubr1("open-dribble-file", ec.openDribbleFile_autogen, 1)
	ec.defSubr0("discard-input", ec.discardInput_autogen)
	ec.defSubr1("suspend-emacs", ec.suspendEmacs_autogen, 0)
	ec.defSubr1("set-input-interrupt-mode", ec.setInputInterruptMode_autogen, 1)
	ec.defSubr2("set-output-flow-control", ec.setOutputFlowControl_autogen, 1)
	ec.defSubr2("set-input-meta-mode", ec.setInputMetaMode_autogen, 1)
	ec.defSubr1("set-quit-char", ec.setQuitChar_autogen, 1)
	ec.defSubr4("set-input-mode", ec.setInputMode_autogen, 3)
	ec.defSubr0("current-input-mode", ec.currentInputMode_autogen)
	ec.defSubr4("posn-at-x-y", ec.posnAtXY_autogen, 2)
	ec.defSubr2("posn-at-point", ec.posnAtPoint_autogen, 0)
	ec.defSubr1("sqlite-open", ec.sqliteOpen_autogen, 0)
	ec.defSubr1("sqlite-close", ec.sqliteClose_autogen, 1)
	ec.defSubr3("sqlite-execute", ec.sqliteExecute_autogen, 2)
	ec.defSubr4("sqlite-select", ec.sqliteSelect_autogen, 2)
	ec.defSubr1("sqlite-transaction", ec.sqliteTransaction_autogen, 1)
	ec.defSubr1("sqlite-commit", ec.sqliteCommit_autogen, 1)
	ec.defSubr1("sqlite-rollback", ec.sqliteRollback_autogen, 1)
	ec.defSubr2("sqlite-pragma", ec.sqlitePragma_autogen, 2)
	ec.defSubr2("sqlite-load-extension", ec.sqliteLoadExtension_autogen, 2)
	ec.defSubr1("sqlite-next", ec.sqliteNext_autogen, 1)
	ec.defSubr1("sqlite-columns", ec.sqliteColumns_autogen, 1)
	ec.defSubr1("sqlite-more-p", ec.sqliteMoreP_autogen, 1)
	ec.defSubr1("sqlite-finalize", ec.sqliteFinalize_autogen, 1)
	ec.defSubr0("sqlite-version", ec.sqliteVersion_autogen)
	ec.defSubr1("sqlitep", ec.sqlitep_autogen, 1)
	ec.defSubr0("sqlite-available-p", ec.sqliteAvailableP_autogen)
	ec.defSubr0("menu-or-popup-active-p", ec.menuOrPopupActiveP_autogen)
	ec.defSubr0("menu-or-popup-active-p", ec.menuOrPopupActiveP_1_autogen)
	ec.defSubr0("menu-or-popup-active-p", ec.menuOrPopupActiveP_2_autogen)
	ec.defSubr0("menu-or-popup-active-p", ec.menuOrPopupActiveP_3_autogen)
	ec.defSubr1("haiku-menu-bar-open", ec.haikuMenuBarOpen_autogen, 0)
	ec.defSubr2("write-char", ec.writeChar_autogen, 1)
	ec.defSubr2("terpri", ec.terpri_autogen, 0)
	ec.defSubr3("prin1", ec.prin1_autogen, 1)
	ec.defSubr3("prin1-to-string", ec.prin1ToString_autogen, 1)
	ec.defSubr2("princ", ec.princ_autogen, 1)
	ec.defSubr2("print", ec.print_autogen, 1)
	ec.defSubr0("flush-standard-output", ec.flushStandardOutput_autogen)
	ec.defSubr1("external-debugging-output", ec.externalDebuggingOutput_autogen, 1)
	ec.defSubr2("redirect-debugging-output", ec.redirectDebuggingOutput_autogen, 1)
	ec.defSubr1("error-message-string", ec.errorMessageString_autogen, 1)
	ec.defSubr1("print--preprocess", ec.printPreprocess_autogen, 1)
	ec.defSubrU("or", ec.or_autogen, 0)
	ec.defSubrU("and", ec.and_autogen, 0)
	ec.defSubrU("if", ec.if_autogen, 2)
	ec.defSubrU("cond", ec.cond_autogen, 0)
	ec.defSubrU("progn", ec.progn_autogen, 0)
	ec.defSubrU("prog1", ec.prog1_autogen, 1)
	ec.defSubrU("setq", ec.setq_autogen, 0)
	ec.defSubrU("quote", ec.quote_autogen, 1)
	ec.defSubrU("function", ec.function_autogen, 1)
	ec.defSubr3("defvaralias", ec.defvaralias_autogen, 2)
	ec.defSubr1("default-toplevel-value", ec.defaultToplevelValue_autogen, 1)
	ec.defSubr2("set-default-toplevel-value", ec.setDefaultToplevelValue_autogen, 2)
	ec.defSubr2("internal--define-uninitialized-variable", ec.internalDefineUninitializedVariable_autogen, 1)
	ec.defSubrU("defvar", ec.defvar_autogen, 1)
	ec.defSubr3("defvar-1", ec.defvar1_autogen, 2)
	ec.defSubrU("defconst", ec.defconst_autogen, 2)
	ec.defSubr3("defconst-1", ec.defconst1_autogen, 2)
	ec.defSubr1("internal-make-var-non-special", ec.makeVarNonSpecial_autogen, 1)
	ec.defSubrU("let*", ec.letX_autogen, 1)
	ec.defSubrU("let", ec.let_autogen, 1)
	ec.defSubrU("while", ec.while_autogen, 1)
	ec.defSubr3("funcall-with-delayed-message", ec.funcallWithDelayedMessage_autogen, 3)
	ec.defSubr2("macroexpand", ec.macroexpand_autogen, 1)
	ec.defSubrU("catch", ec.catch_autogen, 1)
	ec.defSubr2("throw", ec.throw_autogen, 2)
	ec.defSubrU("unwind-protect", ec.unwindProtect_autogen, 1)
	ec.defSubrU("condition-case", ec.conditionCase_autogen, 2)
	ec.defSubr2("signal", ec.signal_autogen, 2)
	ec.defSubr2("commandp", ec.commandp_autogen, 1)
	ec.defSubr5("autoload", ec.autoload_autogen, 2)
	ec.defSubr3("autoload-do-load", ec.autoloadDoLoad_autogen, 1)
	ec.defSubr2("eval", ec.eval_autogen, 1)
	ec.defSubrM("apply", ec.apply_autogen, 1)
	ec.defSubrM("run-hooks", ec.runHooks_autogen, 0)
	ec.defSubrM("run-hook-with-args", ec.runHookWithArgs_autogen, 1)
	ec.defSubrM("run-hook-with-args-until-success", ec.runHookWithArgsUntilSuccess_autogen, 1)
	ec.defSubrM("run-hook-with-args-until-failure", ec.runHookWithArgsUntilFailure_autogen, 1)
	ec.defSubrM("run-hook-wrapped", ec.runHookWrapped_autogen, 2)
	ec.defSubr1("functionp", ec.functionp_autogen, 1)
	ec.defSubrM("funcall", ec.funcall_autogen, 1)
	ec.defSubr1("func-arity", ec.funcArity_autogen, 1)
	ec.defSubr1("fetch-bytecode", ec.fetchBytecode_autogen, 1)
	ec.defSubr1("special-variable-p", ec.specialVariableP_autogen, 1)
	ec.defSubr2("backtrace-debug", ec.backtraceDebug_autogen, 2)
	ec.defSubr2("mapbacktrace", ec.mapbacktrace_autogen, 1)
	ec.defSubr3("backtrace-frame--internal", ec.backtraceFrameInternal_autogen, 3)
	ec.defSubr1("backtrace--frames-from-thread", ec.backtraceFramesFromThread_autogen, 1)
	ec.defSubr3("backtrace-eval", ec.backtraceEval_autogen, 2)
	ec.defSubr2("backtrace--locals", ec.backtraceLocals_autogen, 1)
	ec.defSubr1("x-menu-bar-open-internal", ec.xMenuBarOpenInternal_autogen, 0)
	ec.defSubr1("x-menu-bar-open-internal", ec.xMenuBarOpenInternal_1_autogen, 0)
	ec.defSubr1("x-menu-bar-open-internal", ec.xMenuBarOpenInternal_2_autogen, 0)
	ec.defSubr2("int86", ec.int86_autogen, 2)
	ec.defSubr2("msdos-memget", ec.dosMemget_autogen, 2)
	ec.defSubr2("msdos-memput", ec.dosMemput_autogen, 2)
	ec.defSubr2("msdos-set-keyboard", ec.msdosSetKeyboard_autogen, 1)
	ec.defSubr0("msdos-mouse-p", ec.msdosMouseP_autogen)
	ec.defSubr0("msdos-mouse-init", ec.msdosMouseInit_autogen)
	ec.defSubr0("msdos-mouse-enable", ec.msdosMouseEnable_autogen)
	ec.defSubr0("msdos-mouse-disable", ec.msdosMouseDisable_autogen)
	ec.defSubr0("insert-startup-screen", ec.insertStartupScreen_autogen)
	ec.defSubr1("file-system-info", ec.fileSystemInfo_autogen, 1)
	ec.defSubr1("file-system-info", ec.fileSystemInfo_1_autogen, 1)
	ec.defSubr1("file-system-info", ec.fileSystemInfo_2_autogen, 1)
	ec.defSubr2("text-properties-at", ec.textPropertiesAt_autogen, 1)
	ec.defSubr3("get-text-property", ec.getTextProperty_autogen, 2)
	ec.defSubr3("get-char-property", ec.getCharProperty_autogen, 2)
	ec.defSubr3("get-char-property-and-overlay", ec.getCharPropertyAndOverlay_autogen, 2)
	ec.defSubr2("next-char-property-change", ec.nextCharPropertyChange_autogen, 1)
	ec.defSubr2("previous-char-property-change", ec.previousCharPropertyChange_autogen, 1)
	ec.defSubr4("next-single-char-property-change", ec.nextSingleCharPropertyChange_autogen, 2)
	ec.defSubr4("previous-single-char-property-change", ec.previousSingleCharPropertyChange_autogen, 2)
	ec.defSubr3("next-property-change", ec.nextPropertyChange_autogen, 1)
	ec.defSubr4("next-single-property-change", ec.nextSinglePropertyChange_autogen, 2)
	ec.defSubr3("previous-property-change", ec.previousPropertyChange_autogen, 1)
	ec.defSubr4("previous-single-property-change", ec.previousSinglePropertyChange_autogen, 2)
	ec.defSubr4("add-text-properties", ec.addTextProperties_autogen, 3)
	ec.defSubr5("put-text-property", ec.putTextProperty_autogen, 4)
	ec.defSubr4("set-text-properties", ec.setTextProperties_autogen, 3)
	ec.defSubr5("add-face-text-property", ec.addFaceTextProperty_autogen, 3)
	ec.defSubr4("remove-text-properties", ec.removeTextProperties_autogen, 3)
	ec.defSubr4("remove-list-of-text-properties", ec.removeListOfTextProperties_autogen, 3)
	ec.defSubr5("text-property-any", ec.textPropertyAny_autogen, 4)
	ec.defSubr5("text-property-not-all", ec.textPropertyNotAll_autogen, 4)
	ec.defSubr1("make-mutex", ec.makeMutex_autogen, 0)
	ec.defSubr1("mutex-lock", ec.mutexLock_autogen, 1)
	ec.defSubr1("mutex-unlock", ec.mutexUnlock_autogen, 1)
	ec.defSubr1("mutex-name", ec.mutexName_autogen, 1)
	ec.defSubr2("make-condition-variable", ec.makeConditionVariable_autogen, 1)
	ec.defSubr1("condition-wait", ec.conditionWait_autogen, 1)
	ec.defSubr2("condition-notify", ec.conditionNotify_autogen, 1)
	ec.defSubr1("condition-mutex", ec.conditionMutex_autogen, 1)
	ec.defSubr1("condition-name", ec.conditionName_autogen, 1)
	ec.defSubr0("thread-yield", ec.threadYield_autogen)
	ec.defSubr2("make-thread", ec.makeThread_autogen, 1)
	ec.defSubr0("current-thread", ec.currentThread_autogen)
	ec.defSubr1("thread-name", ec.threadName_autogen, 1)
	ec.defSubr3("thread-signal", ec.threadSignal_autogen, 3)
	ec.defSubr1("thread-live-p", ec.threadLiveP_autogen, 1)
	ec.defSubr1("thread--blocker", ec.threadBlocker_autogen, 1)
	ec.defSubr1("thread-join", ec.threadJoin_autogen, 1)
	ec.defSubr0("all-threads", ec.allThreads_autogen)
	ec.defSubr1("thread-last-error", ec.threadLastError_autogen, 0)
	ec.defSubr0("font-get-system-normal-font", ec.fontGetSystemNormalFont_autogen)
	ec.defSubr0("font-get-system-normal-font", ec.fontGetSystemNormalFont_1_autogen)
	ec.defSubr0("font-get-system-font", ec.fontGetSystemFont_autogen)
	ec.defSubr0("font-get-system-font", ec.fontGetSystemFont_1_autogen)
	ec.defSubr0("tool-bar-get-system-style", ec.toolBarGetSystemStyle_autogen)
	ec.defSubr2("make-char-table", ec.makeCharTable_autogen, 1)
	ec.defSubr1("char-table-subtype", ec.charTableSubtype_autogen, 1)
	ec.defSubr1("char-table-parent", ec.charTableParent_autogen, 1)
	ec.defSubr2("set-char-table-parent", ec.setCharTableParent_autogen, 2)
	ec.defSubr2("char-table-extra-slot", ec.charTableExtraSlot_autogen, 2)
	ec.defSubr3("set-char-table-extra-slot", ec.setCharTableExtraSlot_autogen, 3)
	ec.defSubr2("char-table-range", ec.charTableRange_autogen, 2)
	ec.defSubr3("set-char-table-range", ec.setCharTableRange_autogen, 3)
	ec.defSubr2("optimize-char-table", ec.optimizeCharTable_autogen, 1)
	ec.defSubr2("map-char-table", ec.mapCharTable_autogen, 2)
	ec.defSubr1("unicode-property-table-internal", ec.unicodePropertyTableInternal_autogen, 1)
	ec.defSubr2("get-unicode-property-internal", ec.getUnicodePropertyInternal_autogen, 2)
	ec.defSubr3("put-unicode-property-internal", ec.putUnicodePropertyInternal_autogen, 3)
	ec.defSubr2("x-select-font", ec.xSelectFont_autogen, 0)
	ec.defSubr2("x-select-font", ec.xSelectFont_1_autogen, 0)
	ec.defSubr2("x-select-font", ec.xSelectFont_2_autogen, 0)
	ec.defSubr2("x-select-font", ec.xSelectFont_3_autogen, 0)
	ec.defSubr1("identity", ec.identity_autogen, 1)
	ec.defSubr1("random", ec.random_autogen, 0)
	ec.defSubr1("length", ec.length_autogen, 1)
	ec.defSubr1("safe-length", ec.safeLength_autogen, 1)
	ec.defSubr2("length<", ec.lengthLess_autogen, 2)
	ec.defSubr2("length>", ec.lengthGreater_autogen, 2)
	ec.defSubr2("length=", ec.lengthEqual_autogen, 2)
	ec.defSubr1("proper-list-p", ec.properListP_autogen, 1)
	ec.defSubr1("string-bytes", ec.stringBytes_autogen, 1)
	ec.defSubr3("string-distance", ec.stringDistance_autogen, 2)
	ec.defSubr2("string-equal", ec.stringEqual_autogen, 2)
	ec.defSubr7("compare-strings", ec.compareStrings_autogen, 6)
	ec.defSubr2("string-lessp", ec.stringLessp_autogen, 2)
	ec.defSubr2("string-version-lessp", ec.stringVersionLessp_autogen, 2)
	ec.defSubr4("string-collate-lessp", ec.stringCollateLessp_autogen, 2)
	ec.defSubr4("string-collate-equalp", ec.stringCollateEqualp_autogen, 2)
	ec.defSubrM("append", ec.append_autogen, 0)
	ec.defSubrM("concat", ec.concat_autogen, 0)
	ec.defSubrM("vconcat", ec.vconcat_autogen, 0)
	ec.defSubr1("copy-sequence", ec.copySequence_autogen, 1)
	ec.defSubr1("string-make-multibyte", ec.stringMakeMultibyte_autogen, 1)
	ec.defSubr1("string-make-unibyte", ec.stringMakeUnibyte_autogen, 1)
	ec.defSubr1("string-as-unibyte", ec.stringAsUnibyte_autogen, 1)
	ec.defSubr1("string-as-multibyte", ec.stringAsMultibyte_autogen, 1)
	ec.defSubr1("string-to-multibyte", ec.stringToMultibyte_autogen, 1)
	ec.defSubr1("string-to-unibyte", ec.stringToUnibyte_autogen, 1)
	ec.defSubr1("copy-alist", ec.copyAlist_autogen, 1)
	ec.defSubr3("substring", ec.substring_autogen, 1)
	ec.defSubr3("substring-no-properties", ec.substringNoProperties_autogen, 1)
	ec.defSubr2("take", ec.take_autogen, 2)
	ec.defSubr2("ntake", ec.ntake_autogen, 2)
	ec.defSubr2("nthcdr", ec.nthcdr_autogen, 2)
	ec.defSubr2("nth", ec.nth_autogen, 2)
	ec.defSubr2("elt", ec.elt_autogen, 2)
	ec.defSubr2("member", ec.member_autogen, 2)
	ec.defSubr2("memq", ec.memq_autogen, 2)
	ec.defSubr2("memql", ec.memql_autogen, 2)
	ec.defSubr2("assq", ec.assq_autogen, 2)
	ec.defSubr3("assoc", ec.assoc_autogen, 2)
	ec.defSubr2("rassq", ec.rassq_autogen, 2)
	ec.defSubr2("rassoc", ec.rassoc_autogen, 2)
	ec.defSubr2("delq", ec.delq_autogen, 2)
	ec.defSubr2("delete", ec.delete_autogen, 2)
	ec.defSubr1("nreverse", ec.nreverse_autogen, 1)
	ec.defSubr1("reverse", ec.reverse_autogen, 1)
	ec.defSubr2("sort", ec.sort_autogen, 2)
	ec.defSubr3("plist-get", ec.plistGet_autogen, 2)
	ec.defSubr2("get", ec.get_autogen, 2)
	ec.defSubr4("plist-put", ec.plistPut_autogen, 3)
	ec.defSubr3("put", ec.put_autogen, 3)
	ec.defSubr3("plist-member", ec.plistMember_autogen, 2)
	ec.defSubr2("eql", ec.eql_autogen, 2)
	ec.defSubr2("equal", ec.equal_autogen, 2)
	ec.defSubr2("equal-including-properties", ec.equalIncludingProperties_autogen, 2)
	ec.defSubr2("fillarray", ec.fillarray_autogen, 2)
	ec.defSubr1("clear-string", ec.clearString_autogen, 1)
	ec.defSubrM("nconc", ec.nconc_autogen, 0)
	ec.defSubr3("mapconcat", ec.mapconcat_autogen, 2)
	ec.defSubr2("mapcar", ec.mapcar_autogen, 2)
	ec.defSubr2("mapc", ec.mapc_autogen, 2)
	ec.defSubr2("mapcan", ec.mapcan_autogen, 2)
	ec.defSubr1("yes-or-no-p", ec.yesOrNoP_autogen, 1)
	ec.defSubr1("load-average", ec.loadAverage_autogen, 0)
	ec.defSubr2("featurep", ec.featurep_autogen, 1)
	ec.defSubr2("provide", ec.provide_autogen, 1)
	ec.defSubr3("require", ec.require_autogen, 1)
	ec.defSubr3("widget-put", ec.widgetPut_autogen, 3)
	ec.defSubr2("widget-get", ec.widgetGet_autogen, 2)
	ec.defSubrM("widget-apply", ec.widgetApply_autogen, 2)
	ec.defSubr1("locale-info", ec.localeInfo_autogen, 1)
	ec.defSubr3("base64-encode-region", ec.base64EncodeRegion_autogen, 2)
	ec.defSubr3("base64url-encode-region", ec.base64urlEncodeRegion_autogen, 2)
	ec.defSubr2("base64-encode-string", ec.base64EncodeString_autogen, 1)
	ec.defSubr2("base64url-encode-string", ec.base64urlEncodeString_autogen, 1)
	ec.defSubr4("base64-decode-region", ec.base64DecodeRegion_autogen, 2)
	ec.defSubr3("base64-decode-string", ec.base64DecodeString_autogen, 1)
	ec.defSubr1("sxhash-eq", ec.sxhashEq_autogen, 1)
	ec.defSubr1("sxhash-eql", ec.sxhashEql_autogen, 1)
	ec.defSubr1("sxhash-equal", ec.sxhashEqual_autogen, 1)
	ec.defSubr1("sxhash-equal-including-properties", ec.sxhashEqualIncludingProperties_autogen, 1)
	ec.defSubrM("make-hash-table", ec.makeHashTable_autogen, 0)
	ec.defSubr1("copy-hash-table", ec.copyHashTable_autogen, 1)
	ec.defSubr1("hash-table-count", ec.hashTableCount_autogen, 1)
	ec.defSubr1("hash-table-rehash-size", ec.hashTableRehashSize_autogen, 1)
	ec.defSubr1("hash-table-rehash-threshold", ec.hashTableRehashThreshold_autogen, 1)
	ec.defSubr1("hash-table-size", ec.hashTableSize_autogen, 1)
	ec.defSubr1("hash-table-test", ec.hashTableTest_autogen, 1)
	ec.defSubr1("hash-table-weakness", ec.hashTableWeakness_autogen, 1)
	ec.defSubr1("hash-table-p", ec.hashTableP_autogen, 1)
	ec.defSubr1("clrhash", ec.clrhash_autogen, 1)
	ec.defSubr3("gethash", ec.gethash_autogen, 2)
	ec.defSubr3("puthash", ec.puthash_autogen, 3)
	ec.defSubr2("remhash", ec.remhash_autogen, 2)
	ec.defSubr2("maphash", ec.maphash_autogen, 2)
	ec.defSubr3("define-hash-table-test", ec.defineHashTableTest_autogen, 3)
	ec.defSubr0("secure-hash-algorithms", ec.secureHashAlgorithms_autogen)
	ec.defSubr5("md5", ec.md5_autogen, 1)
	ec.defSubr5("secure-hash", ec.secureHash_autogen, 2)
	ec.defSubr1("buffer-hash", ec.bufferHash_autogen, 0)
	ec.defSubr1("buffer-line-statistics", ec.bufferLineStatistics_autogen, 0)
	ec.defSubr3("string-search", ec.stringSearch_autogen, 2)
	ec.defSubr1("object-intervals", ec.objectIntervals_autogen, 1)
	ec.defSubr2("line-number-at-pos", ec.lineNumberAtPos_autogen, 0)
	ec.defSubr1("handle-save-session", ec.handleSaveSession_autogen, 1)
	ec.defSubr4("set-buffer-redisplay", ec.setBufferRedisplay_autogen, 4)
	ec.defSubr0("line-pixel-height", ec.linePixelHeight_autogen)
	ec.defSubr4("get-display-property", ec.getDisplayProperty_autogen, 2)
	ec.defSubr7("window-text-pixel-size", ec.windowTextPixelSize_autogen, 0)
	ec.defSubr4("buffer-text-pixel-size", ec.bufferTextPixelSize_autogen, 0)
	ec.defSubr0("display--line-is-continued-p", ec.displayLineIsContinuedP_autogen)
	ec.defSubr2("tab-bar-height", ec.tabBarHeight_autogen, 0)
	ec.defSubr2("tool-bar-height", ec.toolBarHeight_autogen, 0)
	ec.defSubr0("long-line-optimizations-p", ec.longLineOptimizationsP_autogen)
	ec.defSubr1("dump-glyph-matrix", ec.dumpGlyphMatrix_autogen, 0)
	ec.defSubr0("dump-frame-glyph-matrix", ec.dumpFrameGlyphMatrix_autogen)
	ec.defSubr2("dump-glyph-row", ec.dumpGlyphRow_autogen, 1)
	ec.defSubr2("dump-tab-bar-row", ec.dumpTabBarRow_autogen, 1)
	ec.defSubr2("dump-tool-bar-row", ec.dumpToolBarRow_autogen, 1)
	ec.defSubr1("trace-redisplay", ec.traceRedisplay_autogen, 0)
	ec.defSubrM("trace-to-stderr", ec.traceToStderr_autogen, 1)
	ec.defSubr1("current-bidi-paragraph-direction", ec.currentBidiParagraphDirection_autogen, 0)
	ec.defSubr4("bidi-find-overridden-directionality", ec.bidiFindOverriddenDirectionality_autogen, 3)
	ec.defSubr1("move-point-visually", ec.movePointVisually_autogen, 1)
	ec.defSubr1("bidi-resolved-levels", ec.bidiResolvedLevels_autogen, 0)
	ec.defSubr4("format-mode-line", ec.formatModeLine_autogen, 1)
	ec.defSubr1("invisible-p", ec.invisibleP_autogen, 1)
	ec.defSubr3("lookup-image-map", ec.lookupImageMap_autogen, 3)
	ec.defSubr2("documentation", ec.documentation_autogen, 1)
	ec.defSubr3("documentation-property", ec.documentationProperty_autogen, 2)
	ec.defSubr1("Snarf-documentation", ec.snarfDocumentation_autogen, 1)
	ec.defSubr0("text-quoting-style", ec.textQuotingStyle_autogen)
	ec.defSubr3("byte-code", ec.byteCode_autogen, 3)
	ec.defSubr0("internal-stack-stats", ec.internalStackStats_autogen)
	ec.defSubr2("query-fontset", ec.queryFontset_autogen, 1)
	ec.defSubr5("set-fontset-font", ec.setFontsetFont_autogen, 3)
	ec.defSubr2("new-fontset", ec.newFontset_autogen, 2)
	ec.defSubr2("fontset-info", ec.fontsetInfo_autogen, 1)
	ec.defSubr3("fontset-font", ec.fontsetFont_autogen, 2)
	ec.defSubr0("fontset-list", ec.fontsetList_autogen)
	ec.defSubr0("fontset-list-all", ec.fontsetListAll_autogen)
	ec.defSubr1("buffer-live-p", ec.bufferLiveP_autogen, 1)
	ec.defSubr1("buffer-list", ec.bufferList_autogen, 0)
	ec.defSubr1("get-buffer", ec.getBuffer_autogen, 1)
	ec.defSubr1("get-file-buffer", ec.getFileBuffer_autogen, 1)
	ec.defSubr2("get-buffer-create", ec.getBufferCreate_autogen, 1)
	ec.defSubr4("make-indirect-buffer", ec.makeIndirectBuffer_autogen, 2)
	ec.defSubr2("generate-new-buffer-name", ec.generateNewBufferName_autogen, 1)
	ec.defSubr1("buffer-name", ec.bufferName_autogen, 0)
	ec.defSubr1("buffer-file-name", ec.bufferFileName_autogen, 0)
	ec.defSubr1("buffer-base-buffer", ec.bufferBaseBuffer_autogen, 0)
	ec.defSubr2("buffer-local-value", ec.bufferLocalValue_autogen, 2)
	ec.defSubr1("buffer-local-variables", ec.bufferLocalVariables_autogen, 0)
	ec.defSubr1("buffer-modified-p", ec.bufferModifiedP_autogen, 0)
	ec.defSubr1("force-mode-line-update", ec.forceModeLineUpdate_autogen, 0)
	ec.defSubr1("set-buffer-modified-p", ec.setBufferModifiedP_autogen, 1)
	ec.defSubr1("restore-buffer-modified-p", ec.restoreBufferModifiedP_autogen, 1)
	ec.defSubr1("buffer-modified-tick", ec.bufferModifiedTick_autogen, 0)
	ec.defSubr2("internal--set-buffer-modified-tick", ec.internalSetBufferModifiedTick_autogen, 1)
	ec.defSubr1("buffer-chars-modified-tick", ec.bufferCharsModifiedTick_autogen, 0)
	ec.defSubr2("rename-buffer", ec.renameBuffer_autogen, 1)
	ec.defSubr3("other-buffer", ec.otherBuffer_autogen, 0)
	ec.defSubr1("buffer-enable-undo", ec.bufferEnableUndo_autogen, 0)
	ec.defSubr1("kill-buffer", ec.killBuffer_autogen, 0)
	ec.defSubr1("bury-buffer-internal", ec.buryBufferInternal_autogen, 1)
	ec.defSubr1("set-buffer-major-mode", ec.setBufferMajorMode_autogen, 1)
	ec.defSubr0("current-buffer", ec.currentBuffer_autogen)
	ec.defSubr1("set-buffer", ec.setBuffer_autogen, 1)
	ec.defSubr1("barf-if-buffer-read-only", ec.barfIfBufferReadOnly_autogen, 0)
	ec.defSubr0("erase-buffer", ec.eraseBuffer_autogen)
	ec.defSubr1("buffer-swap-text", ec.bufferSwapText_autogen, 1)
	ec.defSubr1("set-buffer-multibyte", ec.setBufferMultibyte_autogen, 1)
	ec.defSubr1("kill-all-local-variables", ec.killAllLocalVariables_autogen, 0)
	ec.defSubr1("overlayp", ec.overlayp_autogen, 1)
	ec.defSubr5("make-overlay", ec.makeOverlay_autogen, 2)
	ec.defSubr4("move-overlay", ec.moveOverlay_autogen, 3)
	ec.defSubr1("delete-overlay", ec.deleteOverlay_autogen, 1)
	ec.defSubr1("delete-all-overlays", ec.deleteAllOverlays_autogen, 0)
	ec.defSubr1("overlay-start", ec.overlayStart_autogen, 1)
	ec.defSubr1("overlay-end", ec.overlayEnd_autogen, 1)
	ec.defSubr1("overlay-buffer", ec.overlayBuffer_autogen, 1)
	ec.defSubr1("overlay-properties", ec.overlayProperties_autogen, 1)
	ec.defSubr2("overlays-at", ec.overlaysAt_autogen, 1)
	ec.defSubr2("overlays-in", ec.overlaysIn_autogen, 2)
	ec.defSubr1("next-overlay-change", ec.nextOverlayChange_autogen, 1)
	ec.defSubr1("previous-overlay-change", ec.previousOverlayChange_autogen, 1)
	ec.defSubr0("overlay-lists", ec.overlayLists_autogen)
	ec.defSubr1("overlay-recenter", ec.overlayRecenter_autogen, 1)
	ec.defSubr2("overlay-get", ec.overlayGet_autogen, 2)
	ec.defSubr3("overlay-put", ec.overlayPut_autogen, 3)
	ec.defSubr1("overlay-tree", ec.overlayTree_autogen, 0)
	ec.defSubr1("forward-char", ec.forwardChar_autogen, 0)
	ec.defSubr1("backward-char", ec.backwardChar_autogen, 0)
	ec.defSubr1("forward-line", ec.forwardLine_autogen, 0)
	ec.defSubr1("beginning-of-line", ec.beginningOfLine_autogen, 0)
	ec.defSubr1("end-of-line", ec.endOfLine_autogen, 0)
	ec.defSubr2("delete-char", ec.deleteChar_autogen, 1)
	ec.defSubr2("self-insert-command", ec.selfInsertCommand_autogen, 1)
	ec.defSubr3("read-char", ec.readChar_autogen, 0)
	ec.defSubr3("read-event", ec.readEvent_autogen, 0)
	ec.defSubr3("read-char-exclusive", ec.readCharExclusive_autogen, 0)
	ec.defSubr0("get-file-char", ec.getFileChar_autogen)
	ec.defSubr0("get-load-suffixes", ec.getLoadSuffixes_autogen)
	ec.defSubr5("load", ec.load_autogen, 1)
	ec.defSubr4("locate-file-internal", ec.locateFileInternal_autogen, 2)
	ec.defSubr5("eval-buffer", ec.evalBuffer_autogen, 0)
	ec.defSubr4("eval-region", ec.evalRegion_autogen, 2)
	ec.defSubr1("read", ec.read_autogen, 0)
	ec.defSubr1("read-positioning-symbols", ec.readPositioningSymbols_autogen, 0)
	ec.defSubr3("read-from-string", ec.readFromString_autogen, 1)
	ec.defSubr3("lread--substitute-object-in-subtree", ec.lreadSubstituteObjectInSubtree_autogen, 3)
	ec.defSubr2("intern", ec.intern_autogen, 1)
	ec.defSubr2("intern-soft", ec.internSoft_autogen, 1)
	ec.defSubr2("unintern", ec.unintern_autogen, 1)
	ec.defSubr2("mapatoms", ec.mapatoms_autogen, 1)
	ec.defSubr2("pgtk-use-im-context", ec.pgtkUseImContext_autogen, 1)
	ec.defSubr0("get-internal-run-time", ec.getInternalRunTime_autogen)
	ec.defSubrM("call-process", ec.callProcess_autogen, 1)
	ec.defSubrM("call-process-region", ec.callProcessRegion_autogen, 3)
	ec.defSubr2("getenv-internal", ec.getenvInternal_autogen, 1)
	ec.defSubr2("find-file-name-handler", ec.findFileNameHandler_autogen, 2)
	ec.defSubr1("file-name-directory", ec.fileNameDirectory_autogen, 1)
	ec.defSubr1("file-name-nondirectory", ec.fileNameNondirectory_autogen, 1)
	ec.defSubr1("unhandled-file-name-directory", ec.unhandledFileNameDirectory_autogen, 1)
	ec.defSubr1("file-name-as-directory", ec.fileNameAsDirectory_autogen, 1)
	ec.defSubr1("directory-name-p", ec.directoryNameP_autogen, 1)
	ec.defSubr1("directory-file-name", ec.directoryFileName_autogen, 1)
	ec.defSubr4("make-temp-file-internal", ec.makeTempFileInternal_autogen, 4)
	ec.defSubr1("make-temp-name", ec.makeTempName_autogen, 1)
	ec.defSubrM("file-name-concat", ec.fileNameConcat_autogen, 1)
	ec.defSubr2("expand-file-name", ec.expandFileName_autogen, 1)
	ec.defSubr1("substitute-in-file-name", ec.substituteInFileName_autogen, 1)
	ec.defSubr6("copy-file", ec.copyFile_autogen, 2)
	ec.defSubr1("make-directory-internal", ec.makeDirectoryInternal_autogen, 1)
	ec.defSubr1("delete-directory-internal", ec.deleteDirectoryInternal_autogen, 1)
	ec.defSubr2("delete-file", ec.deleteFile_autogen, 1)
	ec.defSubr1("file-name-case-insensitive-p", ec.fileNameCaseInsensitiveP_autogen, 1)
	ec.defSubr3("rename-file", ec.renameFile_autogen, 2)
	ec.defSubr3("add-name-to-file", ec.addNameToFile_autogen, 2)
	ec.defSubr3("make-symbolic-link", ec.makeSymbolicLink_autogen, 2)
	ec.defSubr1("file-name-absolute-p", ec.fileNameAbsoluteP_autogen, 1)
	ec.defSubr1("file-exists-p", ec.fileExistsP_autogen, 1)
	ec.defSubr1("file-executable-p", ec.fileExecutableP_autogen, 1)
	ec.defSubr1("file-readable-p", ec.fileReadableP_autogen, 1)
	ec.defSubr1("file-writable-p", ec.fileWritableP_autogen, 1)
	ec.defSubr2("access-file", ec.accessFile_autogen, 2)
	ec.defSubr1("file-symlink-p", ec.fileSymlinkP_autogen, 1)
	ec.defSubr1("file-directory-p", ec.fileDirectoryP_autogen, 1)
	ec.defSubr1("file-accessible-directory-p", ec.fileAccessibleDirectoryP_autogen, 1)
	ec.defSubr1("file-regular-p", ec.fileRegularP_autogen, 1)
	ec.defSubr1("file-selinux-context", ec.fileSelinuxContext_autogen, 1)
	ec.defSubr2("set-file-selinux-context", ec.setFileSelinuxContext_autogen, 2)
	ec.defSubr1("file-acl", ec.fileAcl_autogen, 1)
	ec.defSubr2("set-file-acl", ec.setFileAcl_autogen, 2)
	ec.defSubr2("file-modes", ec.fileModes_autogen, 1)
	ec.defSubr3("set-file-modes", ec.setFileModes_autogen, 2)
	ec.defSubr1("set-default-file-modes", ec.setDefaultFileModes_autogen, 1)
	ec.defSubr0("default-file-modes", ec.defaultFileModes_autogen)
	ec.defSubr3("set-file-times", ec.setFileTimes_autogen, 1)
	ec.defSubr0("unix-sync", ec.unixSync_autogen)
	ec.defSubr2("file-newer-than-file-p", ec.fileNewerThanFileP_autogen, 2)
	ec.defSubr5("insert-file-contents", ec.insertFileContents_autogen, 1)
	ec.defSubr7("write-region", ec.writeRegion_autogen, 3)
	ec.defSubr2("car-less-than-car", ec.carLessThanCar_autogen, 2)
	ec.defSubr1("verify-visited-file-modtime", ec.verifyVisitedFileModtime_autogen, 0)
	ec.defSubr0("visited-file-modtime", ec.visitedFileModtime_autogen)
	ec.defSubr1("set-visited-file-modtime", ec.setVisitedFileModtime_autogen, 0)
	ec.defSubr2("do-auto-save", ec.doAutoSave_autogen, 0)
	ec.defSubr0("set-buffer-auto-saved", ec.setBufferAutoSaved_autogen)
	ec.defSubr0("clear-buffer-auto-save-failure", ec.clearBufferAutoSaveFailure_autogen)
	ec.defSubr0("recent-auto-save-p", ec.recentAutoSaveP_autogen)
	ec.defSubr0("next-read-file-uses-dialog-p", ec.nextReadFileUsesDialogP_autogen)
	ec.defSubr2("set-binary-mode", ec.setBinaryMode_autogen, 2)
	ec.defSubr2("haiku-set-mouse-absolute-pixel-position", ec.haikuSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr0("haiku-mouse-absolute-pixel-position", ec.haikuMouseAbsolutePixelPosition_autogen)
	ec.defSubr1("xw-display-color-p", ec.xwDisplayColorP_autogen, 0)
	ec.defSubr1("xw-display-color-p", ec.xwDisplayColorP_1_autogen, 0)
	ec.defSubr1("xw-display-color-p", ec.xwDisplayColorP_2_autogen, 0)
	ec.defSubr1("xw-display-color-p", ec.xwDisplayColorP_3_autogen, 0)
	ec.defSubr2("xw-color-defined-p", ec.xwColorDefinedP_autogen, 1)
	ec.defSubr2("xw-color-defined-p", ec.xwColorDefinedP_1_autogen, 1)
	ec.defSubr2("xw-color-defined-p", ec.xwColorDefinedP_2_autogen, 1)
	ec.defSubr2("xw-color-defined-p", ec.xwColorDefinedP_3_autogen, 1)
	ec.defSubr2("xw-color-values", ec.xwColorValues_autogen, 1)
	ec.defSubr2("xw-color-values", ec.xwColorValues_1_autogen, 1)
	ec.defSubr2("xw-color-values", ec.xwColorValues_2_autogen, 1)
	ec.defSubr2("xw-color-values", ec.xwColorValues_3_autogen, 1)
	ec.defSubr1("x-display-grayscale-p", ec.xDisplayGrayscaleP_autogen, 0)
	ec.defSubr1("x-display-grayscale-p", ec.xDisplayGrayscaleP_1_autogen, 0)
	ec.defSubr1("x-display-grayscale-p", ec.xDisplayGrayscaleP_2_autogen, 0)
	ec.defSubr1("x-display-grayscale-p", ec.xDisplayGrayscaleP_3_autogen, 0)
	ec.defSubr3("x-open-connection", ec.xOpenConnection_autogen, 1)
	ec.defSubr3("x-open-connection", ec.xOpenConnection_1_autogen, 1)
	ec.defSubr3("x-open-connection", ec.xOpenConnection_2_autogen, 1)
	ec.defSubr3("x-open-connection", ec.xOpenConnection_3_autogen, 1)
	ec.defSubr1("x-display-pixel-width", ec.xDisplayPixelWidth_autogen, 0)
	ec.defSubr1("x-display-pixel-width", ec.xDisplayPixelWidth_1_autogen, 0)
	ec.defSubr1("x-display-pixel-width", ec.xDisplayPixelWidth_2_autogen, 0)
	ec.defSubr1("x-display-pixel-width", ec.xDisplayPixelWidth_3_autogen, 0)
	ec.defSubr1("x-display-pixel-height", ec.xDisplayPixelHeight_autogen, 0)
	ec.defSubr1("x-display-pixel-height", ec.xDisplayPixelHeight_1_autogen, 0)
	ec.defSubr1("x-display-pixel-height", ec.xDisplayPixelHeight_2_autogen, 0)
	ec.defSubr1("x-display-pixel-height", ec.xDisplayPixelHeight_3_autogen, 0)
	ec.defSubr1("x-display-mm-height", ec.xDisplayMmHeight_autogen, 0)
	ec.defSubr1("x-display-mm-height", ec.xDisplayMmHeight_1_autogen, 0)
	ec.defSubr1("x-display-mm-height", ec.xDisplayMmHeight_2_autogen, 0)
	ec.defSubr1("x-display-mm-height", ec.xDisplayMmHeight_3_autogen, 0)
	ec.defSubr1("x-display-mm-width", ec.xDisplayMmWidth_autogen, 0)
	ec.defSubr1("x-display-mm-width", ec.xDisplayMmWidth_1_autogen, 0)
	ec.defSubr1("x-display-mm-width", ec.xDisplayMmWidth_2_autogen, 0)
	ec.defSubr1("x-display-mm-width", ec.xDisplayMmWidth_3_autogen, 0)
	ec.defSubr1("x-create-frame", ec.xCreateFrame_autogen, 1)
	ec.defSubr1("x-create-frame", ec.xCreateFrame_1_autogen, 1)
	ec.defSubr1("x-create-frame", ec.xCreateFrame_2_autogen, 1)
	ec.defSubr1("x-create-frame", ec.xCreateFrame_3_autogen, 1)
	ec.defSubr1("x-display-visual-class", ec.xDisplayVisualClass_autogen, 0)
	ec.defSubr1("x-display-visual-class", ec.xDisplayVisualClass_1_autogen, 0)
	ec.defSubr1("x-display-visual-class", ec.xDisplayVisualClass_2_autogen, 0)
	ec.defSubr1("x-display-visual-class", ec.xDisplayVisualClass_3_autogen, 0)
	ec.defSubr6("x-show-tip", ec.xShowTip_autogen, 1)
	ec.defSubr6("x-show-tip", ec.xShowTip_1_autogen, 1)
	ec.defSubr6("x-show-tip", ec.xShowTip_2_autogen, 1)
	ec.defSubr6("x-show-tip", ec.xShowTip_3_autogen, 1)
	ec.defSubr0("x-hide-tip", ec.xHideTip_autogen)
	ec.defSubr0("x-hide-tip", ec.xHideTip_1_autogen)
	ec.defSubr0("x-hide-tip", ec.xHideTip_2_autogen)
	ec.defSubr0("x-hide-tip", ec.xHideTip_3_autogen)
	ec.defSubr1("x-close-connection", ec.xCloseConnection_autogen, 1)
	ec.defSubr1("x-close-connection", ec.xCloseConnection_1_autogen, 1)
	ec.defSubr1("x-close-connection", ec.xCloseConnection_2_autogen, 1)
	ec.defSubr1("x-close-connection", ec.xCloseConnection_3_autogen, 1)
	ec.defSubr0("x-display-list", ec.xDisplayList_autogen)
	ec.defSubr0("x-display-list", ec.xDisplayList_1_autogen)
	ec.defSubr0("x-display-list", ec.xDisplayList_2_autogen)
	ec.defSubr0("x-display-list", ec.xDisplayList_3_autogen)
	ec.defSubr1("x-server-vendor", ec.xServerVendor_autogen, 0)
	ec.defSubr1("x-server-vendor", ec.xServerVendor_1_autogen, 0)
	ec.defSubr1("x-server-vendor", ec.xServerVendor_2_autogen, 0)
	ec.defSubr1("x-server-version", ec.xServerVersion_autogen, 0)
	ec.defSubr1("x-server-version", ec.xServerVersion_1_autogen, 0)
	ec.defSubr1("x-server-version", ec.xServerVersion_2_autogen, 0)
	ec.defSubr1("x-display-screens", ec.xDisplayScreens_autogen, 0)
	ec.defSubr1("x-display-screens", ec.xDisplayScreens_1_autogen, 0)
	ec.defSubr1("x-display-screens", ec.xDisplayScreens_2_autogen, 0)
	ec.defSubr1("x-display-screens", ec.xDisplayScreens_3_autogen, 0)
	ec.defSubr0("haiku-get-version-string", ec.haikuGetVersionString_autogen)
	ec.defSubr1("x-display-color-cells", ec.xDisplayColorCells_autogen, 0)
	ec.defSubr1("x-display-color-cells", ec.xDisplayColorCells_1_autogen, 0)
	ec.defSubr1("x-display-color-cells", ec.xDisplayColorCells_2_autogen, 0)
	ec.defSubr1("x-display-color-cells", ec.xDisplayColorCells_3_autogen, 0)
	ec.defSubr1("x-display-planes", ec.xDisplayPlanes_autogen, 0)
	ec.defSubr1("x-display-planes", ec.xDisplayPlanes_1_autogen, 0)
	ec.defSubr1("x-display-planes", ec.xDisplayPlanes_2_autogen, 0)
	ec.defSubr1("x-display-planes", ec.xDisplayPlanes_3_autogen, 0)
	ec.defSubr1("x-double-buffered-p", ec.xDoubleBufferedP_autogen, 0)
	ec.defSubr1("x-double-buffered-p", ec.xDoubleBufferedP_1_autogen, 0)
	ec.defSubr1("x-display-backing-store", ec.xDisplayBackingStore_autogen, 0)
	ec.defSubr1("x-display-backing-store", ec.xDisplayBackingStore_1_autogen, 0)
	ec.defSubr1("x-display-backing-store", ec.xDisplayBackingStore_2_autogen, 0)
	ec.defSubr1("x-display-backing-store", ec.xDisplayBackingStore_3_autogen, 0)
	ec.defSubr1("haiku-frame-geometry", ec.haikuFrameGeometry_autogen, 0)
	ec.defSubr2("haiku-frame-edges", ec.haikuFrameEdges_autogen, 0)
	ec.defSubr6("haiku-read-file-name", ec.haikuReadFileName_autogen, 1)
	ec.defSubr2("haiku-put-resource", ec.haikuPutResource_autogen, 2)
	ec.defSubr1("haiku-frame-list-z-order", ec.haikuFrameListZOrder_autogen, 0)
	ec.defSubr1("x-display-save-under", ec.xDisplaySaveUnder_autogen, 0)
	ec.defSubr1("x-display-save-under", ec.xDisplaySaveUnder_1_autogen, 0)
	ec.defSubr1("x-display-save-under", ec.xDisplaySaveUnder_2_autogen, 0)
	ec.defSubr1("x-display-save-under", ec.xDisplaySaveUnder_3_autogen, 0)
	ec.defSubr3("haiku-frame-restack", ec.haikuFrameRestack_autogen, 2)
	ec.defSubr1("haiku-save-session-reply", ec.haikuSaveSessionReply_autogen, 1)
	ec.defSubr1("haiku-display-monitor-attributes-list", ec.haikuDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr3("x-own-selection-internal", ec.xOwnSelectionInternal_autogen, 2)
	ec.defSubr4("x-get-selection-internal", ec.xGetSelectionInternal_autogen, 2)
	ec.defSubr3("x-disown-selection-internal", ec.xDisownSelectionInternal_autogen, 1)
	ec.defSubr2("x-selection-owner-p", ec.xSelectionOwnerP_autogen, 0)
	ec.defSubr2("x-selection-exists-p", ec.xSelectionExistsP_autogen, 0)
	ec.defSubr2("x-get-local-selection", ec.xGetLocalSelection_autogen, 0)
	ec.defSubr2("x-get-atom-name", ec.xGetAtomName_autogen, 1)
	ec.defSubr2("x-register-dnd-atom", ec.xRegisterDndAtom_autogen, 1)
	ec.defSubr6("x-send-client-message", ec.xSendClientMessage_autogen, 6)
	ec.defSubr3("w32notify-add-watch", ec.w32notifyAddWatch_autogen, 3)
	ec.defSubr1("w32notify-rm-watch", ec.w32notifyRmWatch_autogen, 1)
	ec.defSubr1("w32notify-valid-p", ec.w32notifyValidP_autogen, 1)
	ec.defSubr2("w32-set-clipboard-data", ec.w32SetClipboardData_autogen, 1)
	ec.defSubr1("w32-get-clipboard-data", ec.w32GetClipboardData_autogen, 0)
	ec.defSubr2("w32-selection-exists-p", ec.w32SelectionExistsP_autogen, 0)
	ec.defSubr2("w32-selection-targets", ec.w32SelectionTargets_autogen, 0)
	ec.defSubr2("gnutls-asynchronous-parameters", ec.gnutlsAsynchronousParameters_autogen, 2)
	ec.defSubr1("gnutls-get-initstage", ec.gnutlsGetInitstage_autogen, 1)
	ec.defSubr1("gnutls-errorp", ec.gnutlsErrorp_autogen, 1)
	ec.defSubr1("gnutls-error-fatalp", ec.gnutlsErrorFatalp_autogen, 1)
	ec.defSubr1("gnutls-error-string", ec.gnutlsErrorString_autogen, 1)
	ec.defSubr1("gnutls-deinit", ec.gnutlsDeinit_autogen, 1)
	ec.defSubr1("gnutls-peer-status-warning-describe", ec.gnutlsPeerStatusWarningDescribe_autogen, 1)
	ec.defSubr1("gnutls-peer-status", ec.gnutlsPeerStatus_autogen, 1)
	ec.defSubr1("gnutls-format-certificate", ec.gnutlsFormatCertificate_autogen, 1)
	ec.defSubr3("gnutls-boot", ec.gnutlsBoot_autogen, 3)
	ec.defSubr2("gnutls-bye", ec.gnutlsBye_autogen, 2)
	ec.defSubr0("gnutls-ciphers", ec.gnutlsCiphers_autogen)
	ec.defSubr5("gnutls-symmetric-encrypt", ec.gnutlsSymmetricEncrypt_autogen, 4)
	ec.defSubr5("gnutls-symmetric-decrypt", ec.gnutlsSymmetricDecrypt_autogen, 4)
	ec.defSubr0("gnutls-macs", ec.gnutlsMacs_autogen)
	ec.defSubr0("gnutls-digests", ec.gnutlsDigests_autogen)
	ec.defSubr3("gnutls-hash-mac", ec.gnutlsHashMac_autogen, 3)
	ec.defSubr2("gnutls-hash-digest", ec.gnutlsHashDigest_autogen, 2)
	ec.defSubr0("gnutls-available-p", ec.gnutlsAvailableP_autogen)
	ec.defSubr5("lcms-cie-de2000", ec.lcmsCieDe2000_autogen, 2)
	ec.defSubr3("lcms-xyz->jch", ec.lcmsXyzToJch_autogen, 1)
	ec.defSubr3("lcms-jch->xyz", ec.lcmsJchToXyz_autogen, 1)
	ec.defSubr3("lcms-jch->jab", ec.lcmsJchToJab_autogen, 1)
	ec.defSubr3("lcms-jab->jch", ec.lcmsJabToJch_autogen, 1)
	ec.defSubr4("lcms-cam02-ucs", ec.lcmsCam02Ucs_autogen, 2)
	ec.defSubr1("lcms-temp->white-point", ec.lcmsTempToWhitePoint_autogen, 1)
	ec.defSubr0("lcms2-available-p", ec.lcms2AvailableP_autogen)
	ec.defSubr1("coding-system-p", ec.codingSystemP_autogen, 1)
	ec.defSubr1("read-non-nil-coding-system", ec.readNonNilCodingSystem_autogen, 1)
	ec.defSubr2("read-coding-system", ec.readCodingSystem_autogen, 1)
	ec.defSubr1("check-coding-system", ec.checkCodingSystem_autogen, 1)
	ec.defSubr3("detect-coding-region", ec.detectCodingRegion_autogen, 2)
	ec.defSubr2("detect-coding-string", ec.detectCodingString_autogen, 1)
	ec.defSubr3("find-coding-systems-region-internal", ec.findCodingSystemsRegionInternal_autogen, 2)
	ec.defSubr5("unencodable-char-position", ec.unencodableCharPosition_autogen, 3)
	ec.defSubr3("check-coding-systems-region", ec.checkCodingSystemsRegion_autogen, 3)
	ec.defSubr4("decode-coding-region", ec.decodeCodingRegion_autogen, 3)
	ec.defSubr4("encode-coding-region", ec.encodeCodingRegion_autogen, 3)
	ec.defSubr7("internal-encode-string-utf-8", ec.internalEncodeStringUtf8_autogen, 7)
	ec.defSubr7("internal-decode-string-utf-8", ec.internalDecodeStringUtf8_autogen, 7)
	ec.defSubr4("decode-coding-string", ec.decodeCodingString_autogen, 2)
	ec.defSubr4("encode-coding-string", ec.encodeCodingString_autogen, 2)
	ec.defSubr1("decode-sjis-char", ec.decodeSjisChar_autogen, 1)
	ec.defSubr1("encode-sjis-char", ec.encodeSjisChar_autogen, 1)
	ec.defSubr1("decode-big5-char", ec.decodeBig5Char_autogen, 1)
	ec.defSubr1("encode-big5-char", ec.encodeBig5Char_autogen, 1)
	ec.defSubr2("set-terminal-coding-system-internal", ec.setTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr1("set-safe-terminal-coding-system-internal", ec.setSafeTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr1("terminal-coding-system", ec.terminalCodingSystem_autogen, 0)
	ec.defSubr2("set-keyboard-coding-system-internal", ec.setKeyboardCodingSystemInternal_autogen, 1)
	ec.defSubr1("keyboard-coding-system", ec.keyboardCodingSystem_autogen, 0)
	ec.defSubrM("find-operation-coding-system", ec.findOperationCodingSystem_autogen, 1)
	ec.defSubrM("set-coding-system-priority", ec.setCodingSystemPriority_autogen, 0)
	ec.defSubr1("coding-system-priority-list", ec.codingSystemPriorityList_autogen, 0)
	ec.defSubrM("define-coding-system-internal", ec.defineCodingSystemInternal_autogen, emacsConstant_coding_arg_max_autogen)
	ec.defSubr3("coding-system-put", ec.codingSystemPut_autogen, 3)
	ec.defSubr2("define-coding-system-alias", ec.defineCodingSystemAlias_autogen, 2)
	ec.defSubr1("coding-system-base", ec.codingSystemBase_autogen, 1)
	ec.defSubr1("coding-system-plist", ec.codingSystemPlist_autogen, 1)
	ec.defSubr1("coding-system-aliases", ec.codingSystemAliases_autogen, 1)
	ec.defSubr1("coding-system-eol-type", ec.codingSystemEolType_autogen, 1)
	ec.defSubr0("dump-colors", ec.dumpColors_autogen)
	ec.defSubr1("clear-face-cache", ec.clearFaceCache_autogen, 0)
	ec.defSubr1("bitmap-spec-p", ec.bitmapSpecP_autogen, 1)
	ec.defSubr1("color-values-from-color-spec", ec.colorValuesFromColorSpec_autogen, 1)
	ec.defSubr2("color-gray-p", ec.colorGrayP_autogen, 1)
	ec.defSubr3("color-supported-p", ec.colorSupportedP_autogen, 1)
	ec.defSubr2("x-family-fonts", ec.xFamilyFonts_autogen, 0)
	ec.defSubr5("x-list-fonts", ec.xListFonts_autogen, 1)
	ec.defSubr2("internal-make-lisp-face", ec.internalMakeLispFace_autogen, 1)
	ec.defSubr2("internal-lisp-face-p", ec.internalLispFaceP_autogen, 1)
	ec.defSubr4("internal-copy-lisp-face", ec.internalCopyLispFace_autogen, 4)
	ec.defSubr4("internal-set-lisp-face-attribute", ec.internalSetLispFaceAttribute_autogen, 3)
	ec.defSubr3("internal-face-x-get-resource", ec.internalFaceXGetResource_autogen, 2)
	ec.defSubr4("internal-set-lisp-face-attribute-from-resource", ec.internalSetLispFaceAttributeFromResource_autogen, 3)
	ec.defSubr2("face-attribute-relative-p", ec.faceAttributeRelativeP_autogen, 2)
	ec.defSubr3("merge-face-attribute", ec.mergeFaceAttribute_autogen, 3)
	ec.defSubr3("internal-get-lisp-face-attribute", ec.internalGetLispFaceAttribute_autogen, 2)
	ec.defSubr1("internal-lisp-face-attribute-values", ec.internalLispFaceAttributeValues_autogen, 1)
	ec.defSubr2("internal-merge-in-global-face", ec.internalMergeInGlobalFace_autogen, 2)
	ec.defSubr3("face-font", ec.faceFont_autogen, 1)
	ec.defSubr3("internal-lisp-face-equal-p", ec.internalLispFaceEqualP_autogen, 2)
	ec.defSubr2("internal-lisp-face-empty-p", ec.internalLispFaceEmptyP_autogen, 1)
	ec.defSubr1("frame--face-hash-table", ec.frameFaceHashTable_autogen, 0)
	ec.defSubr4("color-distance", ec.colorDistance_autogen, 2)
	ec.defSubr1("face-attributes-as-vector", ec.faceAttributesAsVector_autogen, 1)
	ec.defSubr2("display-supports-face-attributes-p", ec.displaySupportsFaceAttributesP_autogen, 1)
	ec.defSubr1("internal-set-font-selection-order", ec.internalSetFontSelectionOrder_autogen, 1)
	ec.defSubr1("internal-set-alternative-font-family-alist", ec.internalSetAlternativeFontFamilyAlist_autogen, 1)
	ec.defSubr1("internal-set-alternative-font-registry-alist", ec.internalSetAlternativeFontRegistryAlist_autogen, 1)
	ec.defSubr1("tty-suppress-bold-inverse-default-colors", ec.ttySuppressBoldInverseDefaultColors_autogen, 1)
	ec.defSubr1("x-load-color-file", ec.xLoadColorFile_autogen, 1)
	ec.defSubr1("dump-face", ec.dumpFace_autogen, 0)
	ec.defSubr0("show-face-resources", ec.showFaceResources_autogen)
	ec.defSubr1("make-keymap", ec.makeKeymap_autogen, 0)
	ec.defSubr1("make-sparse-keymap", ec.makeSparseKeymap_autogen, 0)
	ec.defSubr1("keymapp", ec.keymapp_autogen, 1)
	ec.defSubr1("keymap-prompt", ec.keymapPrompt_autogen, 1)
	ec.defSubr1("keymap-parent", ec.keymapParent_autogen, 1)
	ec.defSubr2("set-keymap-parent", ec.setKeymapParent_autogen, 2)
	ec.defSubr2("map-keymap-internal", ec.mapKeymapInternal_autogen, 2)
	ec.defSubr3("map-keymap", ec.mapKeymap_autogen, 2)
	ec.defSubr2("keymap--get-keyelt", ec.keymapGetKeyelt_autogen, 2)
	ec.defSubr1("copy-keymap", ec.copyKeymap_autogen, 1)
	ec.defSubr4("define-key", ec.defineKey_autogen, 3)
	ec.defSubr3("command-remapping", ec.commandRemapping_autogen, 1)
	ec.defSubr3("lookup-key", ec.lookupKey_autogen, 2)
	ec.defSubr2("current-active-maps", ec.currentActiveMaps_autogen, 0)
	ec.defSubr4("key-binding", ec.keyBinding_autogen, 1)
	ec.defSubr2("minor-mode-key-binding", ec.minorModeKeyBinding_autogen, 1)
	ec.defSubr1("use-global-map", ec.useGlobalMap_autogen, 1)
	ec.defSubr1("use-local-map", ec.useLocalMap_autogen, 1)
	ec.defSubr0("current-local-map", ec.currentLocalMap_autogen)
	ec.defSubr0("current-global-map", ec.currentGlobalMap_autogen)
	ec.defSubr0("current-minor-mode-maps", ec.currentMinorModeMaps_autogen)
	ec.defSubr2("accessible-keymaps", ec.accessibleKeymaps_autogen, 1)
	ec.defSubr2("key-description", ec.keyDescription_autogen, 1)
	ec.defSubr2("single-key-description", ec.singleKeyDescription_autogen, 1)
	ec.defSubr1("text-char-description", ec.textCharDescription_autogen, 1)
	ec.defSubr5("where-is-internal", ec.whereIsInternal_autogen, 1)
	ec.defSubr3("describe-buffer-bindings", ec.describeBufferBindings_autogen, 1)
	ec.defSubr2("describe-vector", ec.describeVector_autogen, 1)
	ec.defSubr7("help--describe-vector", ec.helpDescribeVector_autogen, 7)
	ec.defSubr2("x-export-frames", ec.xExportFrames_autogen, 0)
	ec.defSubr2("x-export-frames", ec.xExportFrames_1_autogen, 0)
	ec.defSubr2("pgtk-set-monitor-scale-factor", ec.pgtkSetMonitorScaleFactor_autogen, 2)
	ec.defSubr3("pgtk-frame-restack", ec.pgtkFrameRestack_autogen, 2)
	ec.defSubr2("pgtk-set-resource", ec.pgtkSetResource_autogen, 2)
	ec.defSubr1("x-server-max-request-size", ec.xServerMaxRequestSize_autogen, 0)
	ec.defSubr1("x-server-max-request-size", ec.xServerMaxRequestSize_1_autogen, 0)
	ec.defSubr1("x-server-max-request-size", ec.xServerMaxRequestSize_2_autogen, 0)
	ec.defSubr1("pgtk-font-name", ec.pgtkFontName_autogen, 1)
	ec.defSubr1("pgtk-display-monitor-attributes-list", ec.pgtkDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr1("pgtk-frame-geometry", ec.pgtkFrameGeometry_autogen, 0)
	ec.defSubr2("pgtk-frame-edges", ec.pgtkFrameEdges_autogen, 0)
	ec.defSubr2("pgtk-set-mouse-absolute-pixel-position", ec.pgtkSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr0("pgtk-mouse-absolute-pixel-position", ec.pgtkMouseAbsolutePixelPosition_autogen)
	ec.defSubr0("pgtk-page-setup-dialog", ec.pgtkPageSetupDialog_autogen)
	ec.defSubr0("pgtk-get-page-setup", ec.pgtkGetPageSetup_autogen)
	ec.defSubr1("pgtk-print-frames-dialog", ec.pgtkPrintFramesDialog_autogen, 0)
	ec.defSubr5("x-file-dialog", ec.xFileDialog_autogen, 2)
	ec.defSubr5("x-file-dialog", ec.xFileDialog_1_autogen, 2)
	ec.defSubr5("x-file-dialog", ec.xFileDialog_2_autogen, 2)
	ec.defSubr5("x-file-dialog", ec.xFileDialog_3_autogen, 2)
	ec.defSubr1("pgtk-backend-display-class", ec.pgtkBackendDisplayClass_autogen, 0)
	ec.defSubr1("x-gtk-debug", ec.xGtkDebug_autogen, 1)
	ec.defSubr1("x-gtk-debug", ec.xGtkDebug_1_autogen, 1)
	ec.defSubr0("current-column", ec.currentColumn_autogen)
	ec.defSubr2("indent-to", ec.indentTo_autogen, 1)
	ec.defSubr0("current-indentation", ec.currentIndentation_autogen)
	ec.defSubr2("move-to-column", ec.moveToColumn_autogen, 1)
	ec.defSubr7("compute-motion", ec.computeMotion_autogen, 7)
	ec.defSubr1("line-number-display-width", ec.lineNumberDisplayWidth_autogen, 0)
	ec.defSubr3("vertical-motion", ec.verticalMotion_autogen, 1)
	ec.defSubr1("w32-has-winsock", ec.w32HasWinsock_autogen, 0)
	ec.defSubr0("w32-unload-winsock", ec.w32UnloadWinsock_autogen)
	ec.defSubr1("w32-short-file-name", ec.w32ShortFileName_autogen, 1)
	ec.defSubr1("w32-long-file-name", ec.w32LongFileName_autogen, 1)
	ec.defSubr2("w32-set-process-priority", ec.w32SetProcessPriority_autogen, 2)
	ec.defSubr1("w32-application-type", ec.w32ApplicationType_autogen, 1)
	ec.defSubr2("w32-get-locale-info", ec.w32GetLocaleInfo_autogen, 1)
	ec.defSubr0("w32-get-current-locale-id", ec.w32GetCurrentLocaleId_autogen)
	ec.defSubr0("w32-get-valid-locale-ids", ec.w32GetValidLocaleIds_autogen)
	ec.defSubr1("w32-get-default-locale-id", ec.w32GetDefaultLocaleId_autogen, 0)
	ec.defSubr1("w32-set-current-locale", ec.w32SetCurrentLocale_autogen, 1)
	ec.defSubr0("w32-get-valid-codepages", ec.w32GetValidCodepages_autogen)
	ec.defSubr0("w32-get-console-codepage", ec.w32GetConsoleCodepage_autogen)
	ec.defSubr1("w32-set-console-codepage", ec.w32SetConsoleCodepage_autogen, 1)
	ec.defSubr0("w32-get-console-output-codepage", ec.w32GetConsoleOutputCodepage_autogen)
	ec.defSubr1("w32-set-console-output-codepage", ec.w32SetConsoleOutputCodepage_autogen, 1)
	ec.defSubr1("w32-get-codepage-charset", ec.w32GetCodepageCharset_autogen, 1)
	ec.defSubr0("w32-get-valid-keyboard-layouts", ec.w32GetValidKeyboardLayouts_autogen)
	ec.defSubr0("w32-get-keyboard-layout", ec.w32GetKeyboardLayout_autogen)
	ec.defSubr1("w32-set-keyboard-layout", ec.w32SetKeyboardLayout_autogen, 1)
	ec.defSubr1("windowp", ec.windowp_autogen, 1)
	ec.defSubr1("window-valid-p", ec.windowValidP_autogen, 1)
	ec.defSubr1("window-live-p", ec.windowLiveP_autogen, 1)
	ec.defSubr1("window-frame", ec.windowFrame_autogen, 0)
	ec.defSubr1("frame-root-window", ec.frameRootWindow_autogen, 0)
	ec.defSubr1("minibuffer-window", ec.minibufferWindow_autogen, 0)
	ec.defSubr1("window-minibuffer-p", ec.windowMinibufferP_autogen, 0)
	ec.defSubr1("frame-first-window", ec.frameFirstWindow_autogen, 0)
	ec.defSubr1("frame-selected-window", ec.frameSelectedWindow_autogen, 0)
	ec.defSubr1("frame-old-selected-window", ec.frameOldSelectedWindow_autogen, 0)
	ec.defSubr3("set-frame-selected-window", ec.setFrameSelectedWindow_autogen, 2)
	ec.defSubr0("selected-window", ec.selectedWindow_autogen)
	ec.defSubr0("old-selected-window", ec.oldSelectedWindow_autogen)
	ec.defSubr2("select-window", ec.selectWindow_autogen, 1)
	ec.defSubr1("window-buffer", ec.windowBuffer_autogen, 0)
	ec.defSubr1("window-old-buffer", ec.windowOldBuffer_autogen, 0)
	ec.defSubr1("window-parent", ec.windowParent_autogen, 0)
	ec.defSubr1("window-top-child", ec.windowTopChild_autogen, 0)
	ec.defSubr1("window-left-child", ec.windowLeftChild_autogen, 0)
	ec.defSubr1("window-next-sibling", ec.windowNextSibling_autogen, 0)
	ec.defSubr1("window-prev-sibling", ec.windowPrevSibling_autogen, 0)
	ec.defSubr1("window-combination-limit", ec.windowCombinationLimit_autogen, 1)
	ec.defSubr2("set-window-combination-limit", ec.setWindowCombinationLimit_autogen, 2)
	ec.defSubr1("window-use-time", ec.windowUseTime_autogen, 0)
	ec.defSubr1("window-bump-use-time", ec.windowBumpUseTime_autogen, 0)
	ec.defSubr1("window-pixel-width", ec.windowPixelWidth_autogen, 0)
	ec.defSubr1("window-pixel-height", ec.windowPixelHeight_autogen, 0)
	ec.defSubr1("window-old-pixel-width", ec.windowOldPixelWidth_autogen, 0)
	ec.defSubr1("window-old-pixel-height", ec.windowOldPixelHeight_autogen, 0)
	ec.defSubr2("window-total-height", ec.windowTotalHeight_autogen, 0)
	ec.defSubr2("window-total-width", ec.windowTotalWidth_autogen, 0)
	ec.defSubr1("window-new-total", ec.windowNewTotal_autogen, 0)
	ec.defSubr2("window-normal-size", ec.windowNormalSize_autogen, 0)
	ec.defSubr1("window-new-normal", ec.windowNewNormal_autogen, 0)
	ec.defSubr1("window-new-pixel", ec.windowNewPixel_autogen, 0)
	ec.defSubr1("window-pixel-left", ec.windowPixelLeft_autogen, 0)
	ec.defSubr1("window-pixel-top", ec.windowPixelTop_autogen, 0)
	ec.defSubr1("window-left-column", ec.windowLeftColumn_autogen, 0)
	ec.defSubr1("window-top-line", ec.windowTopLine_autogen, 0)
	ec.defSubr2("window-body-width", ec.windowBodyWidth_autogen, 0)
	ec.defSubr2("window-body-height", ec.windowBodyHeight_autogen, 0)
	ec.defSubr1("window-old-body-pixel-width", ec.windowOldBodyPixelWidth_autogen, 0)
	ec.defSubr1("window-old-body-pixel-height", ec.windowOldBodyPixelHeight_autogen, 0)
	ec.defSubr1("window-mode-line-height", ec.windowModeLineHeight_autogen, 0)
	ec.defSubr1("window-header-line-height", ec.windowHeaderLineHeight_autogen, 0)
	ec.defSubr1("window-tab-line-height", ec.windowTabLineHeight_autogen, 0)
	ec.defSubr1("window-right-divider-width", ec.windowRightDividerWidth_autogen, 0)
	ec.defSubr1("window-bottom-divider-width", ec.windowBottomDividerWidth_autogen, 0)
	ec.defSubr1("window-scroll-bar-width", ec.windowScrollBarWidth_autogen, 0)
	ec.defSubr1("window-scroll-bar-height", ec.windowScrollBarHeight_autogen, 0)
	ec.defSubr1("window-hscroll", ec.windowHscroll_autogen, 0)
	ec.defSubr2("set-window-hscroll", ec.setWindowHscroll_autogen, 2)
	ec.defSubr2("coordinates-in-window-p", ec.coordinatesInWindowP_autogen, 2)
	ec.defSubr3("window-at", ec.windowAt_autogen, 2)
	ec.defSubr1("window-point", ec.windowPoint_autogen, 0)
	ec.defSubr1("window-old-point", ec.windowOldPoint_autogen, 0)
	ec.defSubr1("window-start", ec.windowStart_autogen, 0)
	ec.defSubr2("window-end", ec.windowEnd_autogen, 0)
	ec.defSubr2("set-window-point", ec.setWindowPoint_autogen, 2)
	ec.defSubr3("set-window-start", ec.setWindowStart_autogen, 2)
	ec.defSubr3("pos-visible-in-window-p", ec.posVisibleInWindowP_autogen, 0)
	ec.defSubr2("window-line-height", ec.windowLineHeight_autogen, 0)
	ec.defSubr6("window-lines-pixel-dimensions", ec.windowLinesPixelDimensions_autogen, 0)
	ec.defSubr1("window-dedicated-p", ec.windowDedicatedP_autogen, 0)
	ec.defSubr2("set-window-dedicated-p", ec.setWindowDedicatedP_autogen, 2)
	ec.defSubr1("window-prev-buffers", ec.windowPrevBuffers_autogen, 0)
	ec.defSubr2("set-window-prev-buffers", ec.setWindowPrevBuffers_autogen, 2)
	ec.defSubr1("window-next-buffers", ec.windowNextBuffers_autogen, 0)
	ec.defSubr2("set-window-next-buffers", ec.setWindowNextBuffers_autogen, 2)
	ec.defSubr1("window-parameters", ec.windowParameters_autogen, 0)
	ec.defSubr2("window-parameter", ec.windowParameter_autogen, 2)
	ec.defSubr3("set-window-parameter", ec.setWindowParameter_autogen, 3)
	ec.defSubr1("window-display-table", ec.windowDisplayTable_autogen, 0)
	ec.defSubr2("set-window-display-table", ec.setWindowDisplayTable_autogen, 2)
	ec.defSubr3("next-window", ec.nextWindow_autogen, 0)
	ec.defSubr3("previous-window", ec.previousWindow_autogen, 0)
	ec.defSubr3("window-list", ec.windowList_autogen, 0)
	ec.defSubr3("window-list-1", ec.windowList1_autogen, 0)
	ec.defSubr2("get-buffer-window", ec.getBufferWindow_autogen, 0)
	ec.defSubr2("delete-other-windows-internal", ec.deleteOtherWindowsInternal_autogen, 0)
	ec.defSubr1("run-window-configuration-change-hook", ec.runWindowConfigurationChangeHook_autogen, 0)
	ec.defSubr1("run-window-scroll-functions", ec.runWindowScrollFunctions_autogen, 0)
	ec.defSubr3("set-window-buffer", ec.setWindowBuffer_autogen, 2)
	ec.defSubr1("force-window-update", ec.forceWindowUpdate_autogen, 0)
	ec.defSubr3("set-window-new-pixel", ec.setWindowNewPixel_autogen, 2)
	ec.defSubr3("set-window-new-total", ec.setWindowNewTotal_autogen, 2)
	ec.defSubr2("set-window-new-normal", ec.setWindowNewNormal_autogen, 1)
	ec.defSubr2("window-resize-apply", ec.windowResizeApply_autogen, 0)
	ec.defSubr2("window-resize-apply-total", ec.windowResizeApplyTotal_autogen, 0)
	ec.defSubr4("split-window-internal", ec.splitWindowInternal_autogen, 4)
	ec.defSubr1("delete-window-internal", ec.deleteWindowInternal_autogen, 1)
	ec.defSubr1("resize-mini-window-internal", ec.resizeMiniWindowInternal_autogen, 1)
	ec.defSubr1("scroll-up", ec.scrollUp_autogen, 0)
	ec.defSubr1("scroll-down", ec.scrollDown_autogen, 0)
	ec.defSubr0("other-window-for-scrolling", ec.otherWindowForScrolling_autogen)
	ec.defSubr2("scroll-left", ec.scrollLeft_autogen, 0)
	ec.defSubr2("scroll-right", ec.scrollRight_autogen, 0)
	ec.defSubr0("minibuffer-selected-window", ec.minibufferSelectedWindow_autogen)
	ec.defSubr2("recenter", ec.recenter_autogen, 0)
	ec.defSubr2("window-text-width", ec.windowTextWidth_autogen, 0)
	ec.defSubr2("window-text-height", ec.windowTextHeight_autogen, 0)
	ec.defSubr1("move-to-window-line", ec.moveToWindowLine_autogen, 1)
	ec.defSubr1("window-configuration-p", ec.windowConfigurationP_autogen, 1)
	ec.defSubr1("window-configuration-frame", ec.windowConfigurationFrame_autogen, 1)
	ec.defSubr3("set-window-configuration", ec.setWindowConfiguration_autogen, 1)
	ec.defSubr1("current-window-configuration", ec.currentWindowConfiguration_autogen, 0)
	ec.defSubr3("set-window-margins", ec.setWindowMargins_autogen, 2)
	ec.defSubr1("window-margins", ec.windowMargins_autogen, 0)
	ec.defSubr5("set-window-fringes", ec.setWindowFringes_autogen, 2)
	ec.defSubr1("window-fringes", ec.windowFringes_autogen, 0)
	ec.defSubr6("set-window-scroll-bars", ec.setWindowScrollBars_autogen, 1)
	ec.defSubr1("window-scroll-bars", ec.windowScrollBars_autogen, 0)
	ec.defSubr2("window-vscroll", ec.windowVscroll_autogen, 0)
	ec.defSubr4("set-window-vscroll", ec.setWindowVscroll_autogen, 2)
	ec.defSubr2("window-configuration-equal-p", ec.windowConfigurationEqualP_autogen, 2)
	ec.defSubr2("characterp", ec.characterp_autogen, 1)
	ec.defSubr1("max-char", ec.maxChar_autogen, 0)
	ec.defSubr1("unibyte-char-to-multibyte", ec.unibyteCharToMultibyte_autogen, 1)
	ec.defSubr1("multibyte-char-to-unibyte", ec.multibyteCharToUnibyte_autogen, 1)
	ec.defSubr1("char-width", ec.charWidth_autogen, 1)
	ec.defSubr3("string-width", ec.stringWidth_autogen, 1)
	ec.defSubrM("string", ec.string_autogen, 0)
	ec.defSubrM("unibyte-string", ec.unibyteString_autogen, 0)
	ec.defSubr1("char-resolve-modifiers", ec.charResolveModifiers_autogen, 1)
	ec.defSubr2("get-byte", ec.getByte_autogen, 0)
	ec.defSubr1("charsetp", ec.charsetp_autogen, 1)
	ec.defSubr5("map-charset-chars", ec.mapCharsetChars_autogen, 2)
	ec.defSubrM("define-charset-internal", ec.defineCharsetInternal_autogen, emacsConstant_charset_arg_max_autogen)
	ec.defSubr2("define-charset-alias", ec.defineCharsetAlias_autogen, 2)
	ec.defSubr1("charset-plist", ec.charsetPlist_autogen, 1)
	ec.defSubr2("set-charset-plist", ec.setCharsetPlist_autogen, 2)
	ec.defSubr3("unify-charset", ec.unifyCharset_autogen, 1)
	ec.defSubr2("get-unused-iso-final-char", ec.getUnusedIsoFinalChar_autogen, 2)
	ec.defSubr4("declare-equiv-charset", ec.declareEquivCharset_autogen, 4)
	ec.defSubr3("find-charset-region", ec.findCharsetRegion_autogen, 2)
	ec.defSubr2("find-charset-string", ec.findCharsetString_autogen, 1)
	ec.defSubr2("decode-char", ec.decodeChar_autogen, 2)
	ec.defSubr2("encode-char", ec.encodeChar_autogen, 2)
	ec.defSubr5("make-char", ec.makeChar_autogen, 1)
	ec.defSubr1("split-char", ec.splitChar_autogen, 1)
	ec.defSubr2("char-charset", ec.charCharset_autogen, 1)
	ec.defSubr1("charset-after", ec.charsetAfter_autogen, 0)
	ec.defSubr3("iso-charset", ec.isoCharset_autogen, 3)
	ec.defSubr0("clear-charset-maps", ec.clearCharsetMaps_autogen)
	ec.defSubr1("charset-priority-list", ec.charsetPriorityList_autogen, 0)
	ec.defSubrM("set-charset-priority", ec.setCharsetPriority_autogen, 1)
	ec.defSubr1("charset-id-internal", ec.charsetIdInternal_autogen, 0)
	ec.defSubr1("sort-charsets", ec.sortCharsets_autogen, 1)
	ec.defSubr3("pgtk-own-selection-internal", ec.pgtkOwnSelectionInternal_autogen, 2)
	ec.defSubr4("pgtk-get-selection-internal", ec.pgtkGetSelectionInternal_autogen, 2)
	ec.defSubr3("pgtk-disown-selection-internal", ec.pgtkDisownSelectionInternal_autogen, 1)
	ec.defSubr2("pgtk-selection-owner-p", ec.pgtkSelectionOwnerP_autogen, 0)
	ec.defSubr2("pgtk-selection-exists-p", ec.pgtkSelectionExistsP_autogen, 0)
	ec.defSubr2("pgtk-register-dnd-targets", ec.pgtkRegisterDndTargets_autogen, 2)
	ec.defSubr3("pgtk-drop-finish", ec.pgtkDropFinish_autogen, 3)
	ec.defSubr2("pgtk-update-drop-status", ec.pgtkUpdateDropStatus_autogen, 2)
	ec.defSubr2("looking-at", ec.lookingAt_autogen, 1)
	ec.defSubr2("posix-looking-at", ec.posixLookingAt_autogen, 1)
	ec.defSubr4("string-match", ec.stringMatch_autogen, 2)
	ec.defSubr4("posix-string-match", ec.posixStringMatch_autogen, 2)
	ec.defSubr4("search-backward", ec.searchBackward_autogen, 1)
	ec.defSubr4("search-forward", ec.searchForward_autogen, 1)
	ec.defSubr4("re-search-backward", ec.reSearchBackward_autogen, 1)
	ec.defSubr4("re-search-forward", ec.reSearchForward_autogen, 1)
	ec.defSubr4("posix-search-backward", ec.posixSearchBackward_autogen, 1)
	ec.defSubr4("posix-search-forward", ec.posixSearchForward_autogen, 1)
	ec.defSubr5("replace-match", ec.replaceMatch_autogen, 1)
	ec.defSubr1("match-beginning", ec.matchBeginning_autogen, 1)
	ec.defSubr1("match-end", ec.matchEnd_autogen, 1)
	ec.defSubr3("match-data", ec.matchData_autogen, 0)
	ec.defSubr2("set-match-data", ec.setMatchData_autogen, 1)
	ec.defSubr1("match-data--translate", ec.matchDataTranslate_autogen, 1)
	ec.defSubr1("regexp-quote", ec.regexpQuote_autogen, 1)
	ec.defSubr1("newline-cache-check", ec.newlineCacheCheck_autogen, 0)
	ec.defSubr2("delete-terminal", ec.deleteTerminal_autogen, 0)
	ec.defSubr1("frame-terminal", ec.frameTerminal_autogen, 0)
	ec.defSubr1("terminal-live-p", ec.terminalLiveP_autogen, 1)
	ec.defSubr0("terminal-list", ec.terminalList_autogen)
	ec.defSubr1("terminal-name", ec.terminalName_autogen, 0)
	ec.defSubr1("terminal-parameters", ec.terminalParameters_autogen, 0)
	ec.defSubr2("terminal-parameter", ec.terminalParameter_autogen, 2)
	ec.defSubr3("set-terminal-parameter", ec.setTerminalParameter_autogen, 3)
	ec.defSubr3("menu-bar-menu-at-x-y", ec.menuBarMenuAtXY_autogen, 2)
	ec.defSubr2("x-popup-menu", ec.xPopupMenu_autogen, 2)
	ec.defSubr3("x-popup-dialog", ec.xPopupDialog_autogen, 2)
	ec.defSubr1("marker-buffer", ec.markerBuffer_autogen, 1)
	ec.defSubr1("marker-position", ec.markerPosition_autogen, 1)
	ec.defSubr3("set-marker", ec.setMarker_autogen, 2)
	ec.defSubr2("copy-marker", ec.copyMarker_autogen, 0)
	ec.defSubr1("marker-insertion-type", ec.markerInsertionType_autogen, 1)
	ec.defSubr2("set-marker-insertion-type", ec.setMarkerInsertionType_autogen, 2)
	ec.defSubr2("w16-set-clipboard-data", ec.w16SetClipboardData_autogen, 1)
	ec.defSubr1("w16-get-clipboard-data", ec.w16GetClipboardData_autogen, 0)
	ec.defSubr2("w16-selection-exists-p", ec.w16SelectionExistsP_autogen, 0)
	ec.defSubr0("dump-redisplay-history", ec.dumpRedisplayHistory_autogen)
	ec.defSubr1("redraw-frame", ec.redrawFrame_autogen, 0)
	ec.defSubr0("redraw-display", ec.redrawDisplay_autogen)
	ec.defSubr2("display--update-for-mouse-movement", ec.displayUpdateForMouseMovement_autogen, 2)
	ec.defSubr1("open-termscript", ec.openTermscript_autogen, 1)
	ec.defSubr2("send-string-to-terminal", ec.sendStringToTerminal_autogen, 1)
	ec.defSubr1("ding", ec.ding_autogen, 0)
	ec.defSubr2("sleep-for", ec.sleepFor_autogen, 1)
	ec.defSubr1("redisplay", ec.redisplay_autogen, 0)
	ec.defSubr1("frame-or-buffer-changed-p", ec.frameOrBufferChangedP_autogen, 0)
	ec.defSubr2("internal-show-cursor", ec.internalShowCursor_autogen, 2)
	ec.defSubr1("internal-show-cursor-p", ec.internalShowCursorP_autogen, 0)
	ec.defSubr5("directory-files", ec.directoryFiles_autogen, 1)
	ec.defSubr6("directory-files-and-attributes", ec.directoryFilesAndAttributes_autogen, 1)
	ec.defSubr3("file-name-completion", ec.fileNameCompletion_autogen, 2)
	ec.defSubr2("file-name-all-completions", ec.fileNameAllCompletions_autogen, 2)
	ec.defSubr2("file-attributes", ec.fileAttributes_autogen, 1)
	ec.defSubr2("file-attributes-lessp", ec.fileAttributesLessp_autogen, 2)
	ec.defSubr0("system-users", ec.systemUsers_autogen)
	ec.defSubr0("system-groups", ec.systemGroups_autogen)
	ec.defSubr0("debug-timer-check", ec.debugTimerCheck_autogen)
	ec.defSubr1("upcase", ec.upcase_autogen, 1)
	ec.defSubr1("downcase", ec.downcase_autogen, 1)
	ec.defSubr1("capitalize", ec.capitalize_autogen, 1)
	ec.defSubr1("upcase-initials", ec.upcaseInitials_autogen, 1)
	ec.defSubr3("upcase-region", ec.upcaseRegion_autogen, 2)
	ec.defSubr3("downcase-region", ec.downcaseRegion_autogen, 2)
	ec.defSubr3("capitalize-region", ec.capitalizeRegion_autogen, 2)
	ec.defSubr3("upcase-initials-region", ec.upcaseInitialsRegion_autogen, 2)
	ec.defSubr1("upcase-word", ec.upcaseWord_autogen, 1)
	ec.defSubr1("downcase-word", ec.downcaseWord_autogen, 1)
	ec.defSubr1("capitalize-word", ec.capitalizeWord_autogen, 1)
	ec.defSubr1("x-wm-set-size-hint", ec.xWmSetSizeHint_autogen, 0)
	ec.defSubr1("x-server-input-extension-version", ec.xServerInputExtensionVersion_autogen, 0)
	ec.defSubr1("x-display-monitor-attributes-list", ec.xDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr1("x-frame-geometry", ec.xFrameGeometry_autogen, 0)
	ec.defSubr2("x-frame-edges", ec.xFrameEdges_autogen, 0)
	ec.defSubr1("x-frame-list-z-order", ec.xFrameListZOrder_autogen, 0)
	ec.defSubr3("x-frame-restack", ec.xFrameRestack_autogen, 2)
	ec.defSubr0("x-mouse-absolute-pixel-position", ec.xMouseAbsolutePixelPosition_autogen)
	ec.defSubr2("x-set-mouse-absolute-pixel-position", ec.xSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr6("x-begin-drag", ec.xBeginDrag_autogen, 1)
	ec.defSubr2("x-synchronize", ec.xSynchronize_autogen, 1)
	ec.defSubr2("x-synchronize", ec.xSynchronize_1_autogen, 1)
	ec.defSubr7("x-change-window-property", ec.xChangeWindowProperty_autogen, 2)
	ec.defSubr6("x-change-window-property", ec.xChangeWindowProperty_1_autogen, 2)
	ec.defSubr3("x-delete-window-property", ec.xDeleteWindowProperty_autogen, 1)
	ec.defSubr2("x-delete-window-property", ec.xDeleteWindowProperty_1_autogen, 1)
	ec.defSubr6("x-window-property", ec.xWindowProperty_autogen, 1)
	ec.defSubr6("x-window-property", ec.xWindowProperty_1_autogen, 1)
	ec.defSubr3("x-window-property-attributes", ec.xWindowPropertyAttributes_autogen, 1)
	ec.defSubr6("x-translate-coordinates", ec.xTranslateCoordinates_autogen, 1)
	ec.defSubr0("x-uses-old-gtk-dialog", ec.xUsesOldGtkDialog_autogen)
	ec.defSubr1("x-backspace-delete-keys-p", ec.xBackspaceDeleteKeysP_autogen, 0)
	ec.defSubr1("x-get-modifier-masks", ec.xGetModifierMasks_autogen, 0)
	ec.defSubr0("x-page-setup-dialog", ec.xPageSetupDialog_autogen)
	ec.defSubr0("x-get-page-setup", ec.xGetPageSetup_autogen)
	ec.defSubr1("x-print-frames-dialog", ec.xPrintFramesDialog_autogen, 0)
	ec.defSubr2("x-display-set-last-user-time", ec.xDisplayLastUserTime_autogen, 1)
	ec.defSubr1("x-internal-focus-input-context", ec.xInternalFocusInputContext_autogen, 1)
	ec.defSubr2("set-screen-color", ec.setScreenColor_autogen, 2)
	ec.defSubr0("get-screen-color", ec.getScreenColor_autogen)
	ec.defSubr1("set-cursor-size", ec.setCursorSize_autogen, 1)
	ec.defSubr1("case-table-p", ec.caseTableP_autogen, 1)
	ec.defSubr0("current-case-table", ec.currentCaseTable_autogen)
	ec.defSubr0("standard-case-table", ec.standardCaseTable_autogen)
	ec.defSubr1("set-case-table", ec.setCaseTable_autogen, 1)
	ec.defSubr1("set-standard-case-table", ec.setStandardCaseTable_autogen, 1)
	ec.defSubr1("haiku-selection-timestamp", ec.haikuSelectionTimestamp_autogen, 1)
	ec.defSubr2("haiku-selection-data", ec.haikuSelectionData_autogen, 2)
	ec.defSubr4("haiku-selection-put", ec.haikuSelectionPut_autogen, 2)
	ec.defSubr1("haiku-selection-owner-p", ec.haikuSelectionOwnerP_autogen, 0)
	ec.defSubr4("haiku-drag-message", ec.haikuDragMessage_autogen, 2)
	ec.defSubr2("haiku-roster-launch", ec.haikuRosterLaunch_autogen, 2)
	ec.defSubr3("haiku-write-node-attribute", ec.haikuWriteNodeAttribute_autogen, 3)
	ec.defSubr2("haiku-send-message", ec.haikuSendMessage_autogen, 2)
	ec.defSubr7("make-xwidget", ec.makeXwidget_autogen, 4)
	ec.defSubr1("xwidget-live-p", ec.xwidgetLiveP_autogen, 1)
	ec.defSubr3("xwidget-perform-lispy-event", ec.xwidgetPerformLispyEvent_autogen, 2)
	ec.defSubr1("get-buffer-xwidgets", ec.getBufferXwidgets_autogen, 1)
	ec.defSubr1("xwidget-webkit-uri", ec.xwidgetWebkitUri_autogen, 1)
	ec.defSubr1("xwidget-webkit-title", ec.xwidgetWebkitTitle_autogen, 1)
	ec.defSubr2("xwidget-webkit-goto-uri", ec.xwidgetWebkitGotoUri_autogen, 2)
	ec.defSubr2("xwidget-webkit-goto-history", ec.xwidgetWebkitGotoHistory_autogen, 2)
	ec.defSubr2("xwidget-webkit-zoom", ec.xwidgetWebkitZoom_autogen, 2)
	ec.defSubr3("xwidget-webkit-execute-script", ec.xwidgetWebkitExecuteScript_autogen, 2)
	ec.defSubr3("xwidget-resize", ec.xwidgetResize_autogen, 3)
	ec.defSubr1("xwidget-size-request", ec.xwidgetSizeRequest_autogen, 1)
	ec.defSubr1("xwidgetp", ec.xwidgetp_autogen, 1)
	ec.defSubr1("xwidget-view-p", ec.xwidgetViewP_autogen, 1)
	ec.defSubr1("xwidget-info", ec.xwidgetInfo_autogen, 1)
	ec.defSubr1("xwidget-view-info", ec.xwidgetViewInfo_autogen, 1)
	ec.defSubr1("xwidget-view-model", ec.xwidgetViewModel_autogen, 1)
	ec.defSubr1("xwidget-view-window", ec.xwidgetViewWindow_autogen, 1)
	ec.defSubr1("delete-xwidget-view", ec.deleteXwidgetView_autogen, 1)
	ec.defSubr2("xwidget-view-lookup", ec.xwidgetViewLookup_autogen, 1)
	ec.defSubr1("xwidget-plist", ec.xwidgetPlist_autogen, 1)
	ec.defSubr1("xwidget-buffer", ec.xwidgetBuffer_autogen, 1)
	ec.defSubr2("set-xwidget-buffer", ec.setXwidgetBuffer_autogen, 2)
	ec.defSubr2("set-xwidget-plist", ec.setXwidgetPlist_autogen, 2)
	ec.defSubr2("set-xwidget-query-on-exit-flag", ec.setXwidgetQueryOnExitFlag_autogen, 2)
	ec.defSubr1("xwidget-query-on-exit-flag", ec.xwidgetQueryOnExitFlag_autogen, 1)
	ec.defSubr5("xwidget-webkit-search", ec.xwidgetWebkitSearch_autogen, 2)
	ec.defSubr1("xwidget-webkit-next-result", ec.xwidgetWebkitNextResult_autogen, 1)
	ec.defSubr1("xwidget-webkit-previous-result", ec.xwidgetWebkitPreviousResult_autogen, 1)
	ec.defSubr1("xwidget-webkit-finish-search", ec.xwidgetWebkitFinishSearch_autogen, 1)
	ec.defSubr1("kill-xwidget", ec.killXwidget_autogen, 1)
	ec.defSubr3("xwidget-webkit-load-html", ec.xwidgetWebkitLoadHtml_autogen, 2)
	ec.defSubr2("xwidget-webkit-back-forward-list", ec.xwidgetWebkitBackForwardList_autogen, 1)
	ec.defSubr1("xwidget-webkit-estimated-load-progress", ec.xwidgetWebkitEstimatedLoadProgress_autogen, 1)
	ec.defSubr2("xwidget-webkit-set-cookie-storage-file", ec.xwidgetWebkitSetCookieStorageFile_autogen, 2)
	ec.defSubr1("xwidget-webkit-stop-loading", ec.xwidgetWebkitStopLoading_autogen, 1)
	ec.defSubr0("zlib-available-p", ec.zlibAvailableP_autogen)
	ec.defSubr3("zlib-decompress-region", ec.zlibDecompressRegion_autogen, 2)
	ec.defSubr4("w32-define-rgb-color", ec.w32DefineRgbColor_autogen, 4)
	ec.defSubr1("w32-display-monitor-attributes-list", ec.w32DisplayMonitorAttributesList_autogen, 0)
	ec.defSubr1("set-message-beep", ec.setMessageBeep_autogen, 1)
	ec.defSubr1("system-move-file-to-trash", ec.systemMoveFileToTrash_autogen, 1)
	ec.defSubr2("w32-send-sys-command", ec.w32SendSysCommand_autogen, 1)
	ec.defSubr4("w32-shell-execute", ec.w32ShellExecute_autogen, 2)
	ec.defSubr1("w32-register-hot-key", ec.w32RegisterHotKey_autogen, 1)
	ec.defSubr1("w32-unregister-hot-key", ec.w32UnregisterHotKey_autogen, 1)
	ec.defSubr0("w32-registered-hot-keys", ec.w32RegisteredHotKeys_autogen)
	ec.defSubr1("w32-reconstruct-hot-key", ec.w32ReconstructHotKey_autogen, 1)
	ec.defSubr2("w32-toggle-lock-key", ec.w32ToggleLockKey_autogen, 1)
	ec.defSubr2("w32-window-exists-p", ec.w32WindowExistsP_autogen, 2)
	ec.defSubr1("w32-frame-geometry", ec.w32FrameGeometry_autogen, 0)
	ec.defSubr2("w32-frame-edges", ec.w32FrameEdges_autogen, 0)
	ec.defSubr1("w32-frame-list-z-order", ec.w32FrameListZOrder_autogen, 0)
	ec.defSubr3("w32-frame-restack", ec.w32FrameRestack_autogen, 2)
	ec.defSubr0("w32-mouse-absolute-pixel-position", ec.w32MouseAbsolutePixelPosition_autogen)
	ec.defSubr2("w32-set-mouse-absolute-pixel-position", ec.w32SetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr0("default-printer-name", ec.defaultPrinterName_autogen)
	ec.defSubr0("w32--menu-bar-in-use", ec.w32MenuBarInUse_autogen)
	ec.defSubrM("w32-notification-notify", ec.w32NotificationNotify_autogen, 0)
	ec.defSubr1("w32-notification-close", ec.w32NotificationClose_autogen, 1)
	ec.defSubr0("w32-get-ime-open-status", ec.w32GetImeOpenStatus_autogen)
	ec.defSubr1("w32-set-ime-open-status", ec.w32SetImeOpenStatus_autogen, 1)
	ec.defSubr3("w32-read-registry", ec.w32ReadRegistry_autogen, 3)
	ec.defSubr1("w32-set-wallpaper", ec.w32SetWallpaper_autogen, 1)
	ec.defSubr1("make-category-set", ec.makeCategorySet_autogen, 1)
	ec.defSubr3("define-category", ec.defineCategory_autogen, 2)
	ec.defSubr2("category-docstring", ec.categoryDocstring_autogen, 1)
	ec.defSubr1("get-unused-category", ec.getUnusedCategory_autogen, 0)
	ec.defSubr1("category-table-p", ec.categoryTableP_autogen, 1)
	ec.defSubr0("category-table", ec.categoryTable_autogen)
	ec.defSubr0("standard-category-table", ec.standardCategoryTable_autogen)
	ec.defSubr1("copy-category-table", ec.copyCategoryTable_autogen, 0)
	ec.defSubr0("make-category-table", ec.makeCategoryTable_autogen)
	ec.defSubr1("set-category-table", ec.setCategoryTable_autogen, 1)
	ec.defSubr1("char-category-set", ec.charCategorySet_autogen, 1)
	ec.defSubr1("category-set-mnemonics", ec.categorySetMnemonics_autogen, 1)
	ec.defSubr4("modify-category-entry", ec.modifyCategoryEntry_autogen, 2)
	ec.defSubr0("json--available-p", ec.jsonAvailableP_autogen)
	ec.defSubrM("json-serialize", ec.jsonSerialize_autogen, 1)
	ec.defSubrM("json-insert", ec.jsonInsert_autogen, 1)
	ec.defSubrM("json-parse-string", ec.jsonParseString_autogen, 1)
	ec.defSubrM("json-parse-buffer", ec.jsonParseBuffer_autogen, 0)
	ec.defSubrU("interactive", ec.interactive_autogen, 0)
	ec.defSubrM("funcall-interactively", ec.funcallInteractively_autogen, 1)
	ec.defSubr3("call-interactively", ec.callInteractively_autogen, 1)
	ec.defSubr1("prefix-numeric-value", ec.prefixNumericValue_autogen, 1)
	ec.defSubr1("processp", ec.processp_autogen, 1)
	ec.defSubr1("get-process", ec.getProcess_autogen, 1)
	ec.defSubr1("delete-process", ec.deleteProcess_autogen, 0)
	ec.defSubr1("process-status", ec.processStatus_autogen, 1)
	ec.defSubr1("process-exit-status", ec.processExitStatus_autogen, 1)
	ec.defSubr1("process-id", ec.processId_autogen, 1)
	ec.defSubr1("process-name", ec.processName_autogen, 1)
	ec.defSubr1("process-command", ec.processCommand_autogen, 1)
	ec.defSubr2("process-tty-name", ec.processTtyName_autogen, 1)
	ec.defSubr2("set-process-buffer", ec.setProcessBuffer_autogen, 2)
	ec.defSubr1("process-buffer", ec.processBuffer_autogen, 1)
	ec.defSubr1("process-mark", ec.processMark_autogen, 1)
	ec.defSubr2("set-process-filter", ec.setProcessFilter_autogen, 2)
	ec.defSubr1("process-filter", ec.processFilter_autogen, 1)
	ec.defSubr2("set-process-sentinel", ec.setProcessSentinel_autogen, 2)
	ec.defSubr1("process-sentinel", ec.processSentinel_autogen, 1)
	ec.defSubr2("set-process-thread", ec.setProcessThread_autogen, 2)
	ec.defSubr1("process-thread", ec.processThread_autogen, 1)
	ec.defSubr3("set-process-window-size", ec.setProcessWindowSize_autogen, 3)
	ec.defSubr2("set-process-inherit-coding-system-flag", ec.setProcessInheritCodingSystemFlag_autogen, 2)
	ec.defSubr2("set-process-query-on-exit-flag", ec.setProcessQueryOnExitFlag_autogen, 2)
	ec.defSubr1("process-query-on-exit-flag", ec.processQueryOnExitFlag_autogen, 1)
	ec.defSubr3("process-contact", ec.processContact_autogen, 1)
	ec.defSubr1("process-plist", ec.processPlist_autogen, 1)
	ec.defSubr2("set-process-plist", ec.setProcessPlist_autogen, 2)
	ec.defSubr1("process-connection", ec.processConnection_autogen, 1)
	ec.defSubr1("process-type", ec.processType_autogen, 1)
	ec.defSubr2("format-network-address", ec.formatNetworkAddress_autogen, 1)
	ec.defSubr0("process-list", ec.processList_autogen)
	ec.defSubrM("make-process", ec.makeProcess_autogen, 0)
	ec.defSubrM("make-pipe-process", ec.makePipeProcess_autogen, 0)
	ec.defSubr1("process-datagram-address", ec.processDatagramAddress_autogen, 1)
	ec.defSubr2("set-process-datagram-address", ec.setProcessDatagramAddress_autogen, 2)
	ec.defSubr4("set-network-process-option", ec.setNetworkProcessOption_autogen, 3)
	ec.defSubrM("serial-process-configure", ec.serialProcessConfigure_autogen, 0)
	ec.defSubrM("make-serial-process", ec.makeSerialProcess_autogen, 0)
	ec.defSubrM("make-network-process", ec.makeNetworkProcess_autogen, 0)
	ec.defSubr2("network-interface-list", ec.networkInterfaceList_autogen, 0)
	ec.defSubr1("network-interface-info", ec.networkInterfaceInfo_autogen, 1)
	ec.defSubr3("network-lookup-address-info", ec.networkLookupAddressInfo_autogen, 1)
	ec.defSubr4("accept-process-output", ec.acceptProcessOutput_autogen, 0)
	ec.defSubr2("internal-default-process-filter", ec.internalDefaultProcessFilter_autogen, 2)
	ec.defSubr3("process-send-region", ec.processSendRegion_autogen, 3)
	ec.defSubr2("process-send-string", ec.processSendString_autogen, 2)
	ec.defSubr1("process-running-child-p", ec.processRunningChildP_autogen, 0)
	ec.defSubr2("internal-default-interrupt-process", ec.internalDefaultInterruptProcess_autogen, 0)
	ec.defSubr2("interrupt-process", ec.interruptProcess_autogen, 0)
	ec.defSubr2("kill-process", ec.killProcess_autogen, 0)
	ec.defSubr2("quit-process", ec.quitProcess_autogen, 0)
	ec.defSubr2("stop-process", ec.stopProcess_autogen, 0)
	ec.defSubr2("continue-process", ec.continueProcess_autogen, 0)
	ec.defSubr3("internal-default-signal-process", ec.internalDefaultSignalProcess_autogen, 2)
	ec.defSubr3("signal-process", ec.signalProcess_autogen, 2)
	ec.defSubr1("process-send-eof", ec.processSendEof_autogen, 0)
	ec.defSubr2("internal-default-process-sentinel", ec.internalDefaultProcessSentinel_autogen, 2)
	ec.defSubr3("set-process-coding-system", ec.setProcessCodingSystem_autogen, 1)
	ec.defSubr1("process-coding-system", ec.processCodingSystem_autogen, 1)
	ec.defSubr1("get-buffer-process", ec.getBufferProcess_autogen, 1)
	ec.defSubr1("process-inherit-coding-system-flag", ec.processInheritCodingSystemFlag_autogen, 1)
	ec.defSubr0("waiting-for-user-input-p", ec.waitingForUserInputP_autogen)
	ec.defSubr0("list-system-processes", ec.listSystemProcesses_autogen)
	ec.defSubr1("process-attributes", ec.processAttributes_autogen, 1)
	ec.defSubr1("num-processors", ec.numProcessors_autogen, 0)
	ec.defSubr0("signal-names", ec.signalNames_autogen)
	ec.defSubr3("make-string", ec.makeString_autogen, 2)
	ec.defSubr2("make-bool-vector", ec.makeBoolVector_autogen, 2)
	ec.defSubrM("bool-vector", ec.boolVector_autogen, 0)
	ec.defSubr2("cons", ec.cons_autogen, 2)
	ec.defSubrM("list", ec.list_autogen, 0)
	ec.defSubr2("make-list", ec.makeList_autogen, 2)
	ec.defSubr3("make-record", ec.makeRecord_autogen, 3)
	ec.defSubrM("record", ec.record_autogen, 1)
	ec.defSubr2("make-vector", ec.makeVector_autogen, 2)
	ec.defSubrM("vector", ec.vector_autogen, 0)
	ec.defSubrM("make-byte-code", ec.makeByteCode_autogen, 4)
	ec.defSubrM("make-closure", ec.makeClosure_autogen, 1)
	ec.defSubr1("make-symbol", ec.makeSymbol_autogen, 1)
	ec.defSubr0("make-marker", ec.makeMarker_autogen)
	ec.defSubr1("make-finalizer", ec.makeFinalizer_autogen, 1)
	ec.defSubr1("purecopy", ec.purecopy_autogen, 1)
	ec.defSubr0("garbage-collect", ec.garbageCollect_autogen)
	ec.defSubr1("garbage-collect-maybe", ec.garbageCollectMaybe_autogen, 1)
	ec.defSubr0("memory-info", ec.memoryInfo_autogen)
	ec.defSubr0("memory-use-counts", ec.memoryUseCounts_autogen)
	ec.defSubr0("malloc-info", ec.mallocInfo_autogen)
	ec.defSubr1("malloc-trim", ec.mallocTrim_autogen, 0)
	ec.defSubr1("suspicious-object", ec.suspiciousObject_autogen, 1)
	ec.defSubr1("acos", ec.acos_autogen, 1)
	ec.defSubr1("asin", ec.asin_autogen, 1)
	ec.defSubr2("atan", ec.atan_autogen, 1)
	ec.defSubr1("cos", ec.cos_autogen, 1)
	ec.defSubr1("sin", ec.sin_autogen, 1)
	ec.defSubr1("tan", ec.tan_autogen, 1)
	ec.defSubr1("isnan", ec.isnan_autogen, 1)
	ec.defSubr2("copysign", ec.copysign_autogen, 2)
	ec.defSubr1("frexp", ec.frexp_autogen, 1)
	ec.defSubr2("ldexp", ec.ldexp_autogen, 2)
	ec.defSubr1("exp", ec.exp_autogen, 1)
	ec.defSubr2("expt", ec.expt_autogen, 2)
	ec.defSubr2("log", ec.log_autogen, 1)
	ec.defSubr1("sqrt", ec.sqrt_autogen, 1)
	ec.defSubr1("abs", ec.abs_autogen, 1)
	ec.defSubr1("float", ec.float_autogen, 1)
	ec.defSubr1("logb", ec.logb_autogen, 1)
	ec.defSubr2("ceiling", ec.ceiling_autogen, 1)
	ec.defSubr2("floor", ec.floor_autogen, 1)
	ec.defSubr2("round", ec.round_autogen, 1)
	ec.defSubr2("truncate", ec.truncate_autogen, 1)
	ec.defSubr1("fceiling", ec.fceiling_autogen, 1)
	ec.defSubr1("ffloor", ec.ffloor_autogen, 1)
	ec.defSubr1("fround", ec.fround_autogen, 1)
	ec.defSubr1("ftruncate", ec.ftruncate_autogen, 1)
	ec.defSubr0("w32-battery-status", ec.w32BatteryStatus_autogen)
	ec.defSubr1("comp--subr-signature", ec.compSubrSignature_autogen, 1)
	ec.defSubr1("comp-el-to-eln-rel-filename", ec.compElToElnRelFilename_autogen, 1)
	ec.defSubr2("comp-el-to-eln-filename", ec.compElToElnFilename_autogen, 1)
	ec.defSubr2("comp--install-trampoline", ec.compInstallTrampoline_autogen, 2)
	ec.defSubr0("comp--init-ctxt", ec.compInitCtxt_autogen)
	ec.defSubr0("comp--release-ctxt", ec.compReleaseCtxt_autogen)
	ec.defSubr0("comp-native-driver-options-effective-p", ec.compNativeDriverOptionsEffectiveP_autogen)
	ec.defSubr0("comp-native-compiler-options-effective-p", ec.compNativeCompilerOptionsEffectiveP_autogen)
	ec.defSubr1("comp--compile-ctxt-to-file", ec.compCompileCtxtToFile_autogen, 1)
	ec.defSubr0("comp-libgccjit-version", ec.compLibgccjitVersion_autogen)
	ec.defSubr7("comp--register-lambda", ec.compRegisterLambda_autogen, 7)
	ec.defSubr7("comp--register-subr", ec.compRegisterSubr_autogen, 7)
	ec.defSubr7("comp--late-register-subr", ec.compLateRegisterSubr_autogen, 7)
	ec.defSubr2("native-elisp-load", ec.nativeElispLoad_autogen, 1)
	ec.defSubr0("native-comp-available-p", ec.nativeCompAvailableP_autogen)
	ec.defSubr2("dump-emacs-portable", ec.dumpEmacsPortable_autogen, 1)
	ec.defSubr2("dump-emacs-portable--sort-predicate", ec.dumpEmacsPortableSortPredicate_autogen, 2)
	ec.defSubr2("dump-emacs-portable--sort-predicate-copied", ec.dumpEmacsPortableSortPredicateCopied_autogen, 2)
	ec.defSubr0("pdumper-stats", ec.pdumperStats_autogen)
	ec.defSubr0("invocation-name", ec.invocationName_autogen)
	ec.defSubr0("invocation-directory", ec.invocationDirectory_autogen)
	ec.defSubr2("kill-emacs", ec.killEmacs_autogen, 0)
	ec.defSubr2("dump-emacs", ec.dumpEmacs_autogen, 2)
	ec.defSubr0("daemonp", ec.daemonp_autogen)
	ec.defSubr0("daemon-initialized", ec.daemonInitialized_autogen)
	ec.defSubr2("treesit-language-available-p", ec.treesitLanguageAvailableP_autogen, 1)
	ec.defSubr1("treesit-library-abi-version", ec.treesitLibraryAbiVersion_autogen, 0)
	ec.defSubr1("treesit-language-abi-version", ec.treesitLanguageAbiVersion_autogen, 0)
	ec.defSubr1("treesit-parser-p", ec.treesitParserP_autogen, 1)
	ec.defSubr1("treesit-node-p", ec.treesitNodeP_autogen, 1)
	ec.defSubr1("treesit-compiled-query-p", ec.treesitCompiledQueryP_autogen, 1)
	ec.defSubr1("treesit-query-p", ec.treesitQueryP_autogen, 1)
	ec.defSubr1("treesit-query-language", ec.treesitQueryLanguage_autogen, 1)
	ec.defSubr1("treesit-node-parser", ec.treesitNodeParser_autogen, 1)
	ec.defSubr3("treesit-parser-create", ec.treesitParserCreate_autogen, 1)
	ec.defSubr1("treesit-parser-delete", ec.treesitParserDelete_autogen, 1)
	ec.defSubr1("treesit-parser-list", ec.treesitParserList_autogen, 0)
	ec.defSubr1("treesit-parser-buffer", ec.treesitParserBuffer_autogen, 1)
	ec.defSubr1("treesit-parser-language", ec.treesitParserLanguage_autogen, 1)
	ec.defSubr1("treesit-parser-root-node", ec.treesitParserRootNode_autogen, 1)
	ec.defSubr2("treesit-parser-set-included-ranges", ec.treesitParserSetIncludedRanges_autogen, 2)
	ec.defSubr1("treesit-parser-included-ranges", ec.treesitParserIncludedRanges_autogen, 1)
	ec.defSubr1("treesit-parser-notifiers", ec.treesitParserNotifiers_autogen, 1)
	ec.defSubr2("treesit-parser-add-notifier", ec.treesitParserAddNotifier_autogen, 2)
	ec.defSubr2("treesit-parser-remove-notifier", ec.treesitParserRemoveNotifier_autogen, 2)
	ec.defSubr1("treesit-node-type", ec.treesitNodeType_autogen, 1)
	ec.defSubr1("treesit-node-start", ec.treesitNodeStart_autogen, 1)
	ec.defSubr1("treesit-node-end", ec.treesitNodeEnd_autogen, 1)
	ec.defSubr1("treesit-node-string", ec.treesitNodeString_autogen, 1)
	ec.defSubr1("treesit-node-parent", ec.treesitNodeParent_autogen, 1)
	ec.defSubr3("treesit-node-child", ec.treesitNodeChild_autogen, 2)
	ec.defSubr2("treesit-node-check", ec.treesitNodeCheck_autogen, 2)
	ec.defSubr2("treesit-node-field-name-for-child", ec.treesitNodeFieldNameForChild_autogen, 2)
	ec.defSubr2("treesit-node-child-count", ec.treesitNodeChildCount_autogen, 1)
	ec.defSubr2("treesit-node-child-by-field-name", ec.treesitNodeChildByFieldName_autogen, 2)
	ec.defSubr2("treesit-node-next-sibling", ec.treesitNodeNextSibling_autogen, 1)
	ec.defSubr2("treesit-node-prev-sibling", ec.treesitNodePrevSibling_autogen, 1)
	ec.defSubr3("treesit-node-first-child-for-pos", ec.treesitNodeFirstChildForPos_autogen, 2)
	ec.defSubr4("treesit-node-descendant-for-range", ec.treesitNodeDescendantForRange_autogen, 3)
	ec.defSubr2("treesit-node-eq", ec.treesitNodeEq_autogen, 2)
	ec.defSubr1("treesit-pattern-expand", ec.treesitPatternExpand_autogen, 1)
	ec.defSubr1("treesit-query-expand", ec.treesitQueryExpand_autogen, 1)
	ec.defSubr3("treesit-query-compile", ec.treesitQueryCompile_autogen, 2)
	ec.defSubr5("treesit-query-capture", ec.treesitQueryCapture_autogen, 2)
	ec.defSubr5("treesit-search-subtree", ec.treesitSearchSubtree_autogen, 2)
	ec.defSubr4("treesit-search-forward", ec.treesitSearchForward_autogen, 2)
	ec.defSubr4("treesit-induce-sparse-tree", ec.treesitInduceSparseTree_autogen, 2)
	ec.defSubr1("treesit-subtree-stat", ec.treesitSubtreeStat_autogen, 1)
	ec.defSubr0("treesit-available-p", ec.treesitAvailableP_autogen)
	ec.defSubr2("fontp", ec.fontp_autogen, 1)
	ec.defSubrM("font-spec", ec.fontSpec_autogen, 0)
	ec.defSubr2("font-get", ec.fontGet_autogen, 2)
	ec.defSubr2("font-face-attributes", ec.fontFaceAttributes_autogen, 1)
	ec.defSubr3("font-put", ec.fontPut_autogen, 3)
	ec.defSubr4("list-fonts", ec.listFonts_autogen, 1)
	ec.defSubr1("font-family-list", ec.fontFamilyList_autogen, 0)
	ec.defSubr2("find-font", ec.findFont_autogen, 1)
	ec.defSubr2("font-xlfd-name", ec.fontXlfdName_autogen, 1)
	ec.defSubr0("clear-font-cache", ec.clearFontCache_autogen)
	ec.defSubr2("font-shape-gstring", ec.fontShapeGstring_autogen, 2)
	ec.defSubr2("font-variation-glyphs", ec.fontVariationGlyphs_autogen, 2)
	ec.defSubr2("internal-char-font", ec.internalCharFont_autogen, 1)
	ec.defSubr6("font-drive-otf", ec.fontDriveOtf_autogen, 6)
	ec.defSubr3("font-otf-alternates", ec.fontOtfAlternates_autogen, 3)
	ec.defSubr3("open-font", ec.openFont_autogen, 1)
	ec.defSubr2("close-font", ec.closeFont_autogen, 1)
	ec.defSubr1("query-font", ec.queryFont_autogen, 1)
	ec.defSubr3("font-has-char-p", ec.fontHasCharP_autogen, 2)
	ec.defSubr4("font-get-glyphs", ec.fontGetGlyphs_autogen, 3)
	ec.defSubr2("font-match-p", ec.fontMatchP_autogen, 2)
	ec.defSubr3("font-at", ec.fontAt_autogen, 1)
	ec.defSubr2("draw-string", ec.drawString_autogen, 2)
	ec.defSubr1("frame-font-cache", ec.frameFontCache_autogen, 0)
	ec.defSubr2("font-info", ec.fontInfo_autogen, 1)
	ec.defSubr1("module-load", ec.moduleLoad_autogen, 1)
	ec.defSubr1("play-sound-internal", ec.playSoundInternal_autogen, 1)
	ec.defSubr3("image-size", ec.imageSize_autogen, 1)
	ec.defSubr2("image-mask-p", ec.imageMaskP_autogen, 1)
	ec.defSubr2("image-metadata", ec.imageMetadata_autogen, 1)
	ec.defSubr2("clear-image-cache", ec.clearImageCache_autogen, 0)
	ec.defSubr2("image-flush", ec.imageFlush_autogen, 1)
	ec.defSubr0("imagemagick-types", ec.imagemagickTypes_autogen)
	ec.defSubr1("imagep", ec.imagep_autogen, 1)
	ec.defSubr1("lookup-image", ec.lookupImage_autogen, 1)
	ec.defSubr1("image-transforms-p", ec.imageTransformsP_autogen, 0)
	ec.defSubr0("image-cache-size", ec.imageCacheSize_autogen)
	ec.defSubr1("init-image-library", ec.initImageLibrary_autogen, 1)
	ec.defSubr3("gfile-add-watch", ec.gfileAddWatch_autogen, 3)
	ec.defSubr1("gfile-rm-watch", ec.gfileRmWatch_autogen, 1)
	ec.defSubr1("gfile-valid-p", ec.gfileValidP_autogen, 1)
	ec.defSubr1("gfile-monitor-name", ec.gfileMonitorName_autogen, 1)
	ec.defSubr2("dbus--init-bus", ec.dbusInitBus_autogen, 1)
	ec.defSubr1("dbus-get-unique-name", ec.dbusGetUniqueName_autogen, 1)
	ec.defSubrM("dbus-message-internal", ec.dbusMessageInternal_autogen, 4)
	ec.defSubr1("destroy-fringe-bitmap", ec.destroyFringeBitmap_autogen, 1)
	ec.defSubr5("define-fringe-bitmap", ec.defineFringeBitmap_autogen, 2)
	ec.defSubr2("set-fringe-bitmap-face", ec.setFringeBitmapFace_autogen, 1)
	ec.defSubr2("fringe-bitmaps-at-pos", ec.fringeBitmapsAtPos_autogen, 0)
	ec.defSubr1("framep", ec.framep_autogen, 1)
	ec.defSubr1("frame-live-p", ec.frameLiveP_autogen, 1)
	ec.defSubr1("window-system", ec.windowSystem_autogen, 0)
	ec.defSubr4("frame-windows-min-size", ec.frameWindowsMinSize_autogen, 4)
	ec.defSubr1("make-terminal-frame", ec.makeTerminalFrame_autogen, 1)
	ec.defSubr2("select-frame", ec.selectFrame_autogen, 1)
	ec.defSubr1("handle-switch-frame", ec.handleSwitchFrame_autogen, 1)
	ec.defSubr0("selected-frame", ec.selectedFrame_autogen)
	ec.defSubr0("old-selected-frame", ec.oldSelectedFrame_autogen)
	ec.defSubr0("frame-list", ec.frameList_autogen)
	ec.defSubr1("frame-parent", ec.frameParent_autogen, 0)
	ec.defSubr2("frame-ancestor-p", ec.frameAncestorP_autogen, 2)
	ec.defSubr2("next-frame", ec.nextFrame_autogen, 0)
	ec.defSubr2("previous-frame", ec.previousFrame_autogen, 0)
	ec.defSubr0("last-nonminibuffer-frame", ec.lastNonminibufFrame_autogen)
	ec.defSubr2("delete-frame", ec.deleteFrame_autogen, 0)
	ec.defSubr0("mouse-position", ec.mousePosition_autogen)
	ec.defSubr0("mouse-pixel-position", ec.mousePixelPosition_autogen)
	ec.defSubr3("set-mouse-position", ec.setMousePosition_autogen, 3)
	ec.defSubr3("set-mouse-pixel-position", ec.setMousePixelPosition_autogen, 3)
	ec.defSubr1("make-frame-visible", ec.makeFrameVisible_autogen, 0)
	ec.defSubr2("make-frame-invisible", ec.makeFrameInvisible_autogen, 0)
	ec.defSubr1("iconify-frame", ec.iconifyFrame_autogen, 0)
	ec.defSubr1("frame-visible-p", ec.frameVisibleP_autogen, 1)
	ec.defSubr0("visible-frame-list", ec.visibleFrameList_autogen)
	ec.defSubr1("raise-frame", ec.raiseFrame_autogen, 0)
	ec.defSubr1("lower-frame", ec.lowerFrame_autogen, 0)
	ec.defSubr2("redirect-frame-focus", ec.redirectFrameFocus_autogen, 1)
	ec.defSubr1("frame-focus", ec.frameFocus_autogen, 0)
	ec.defSubr2("x-focus-frame", ec.xFocusFrame_autogen, 1)
	ec.defSubr2("frame-after-make-frame", ec.frameAfterMakeFrame_autogen, 2)
	ec.defSubr1("frame-parameters", ec.frameParameters_autogen, 0)
	ec.defSubr2("frame-parameter", ec.frameParameter_autogen, 2)
	ec.defSubr2("modify-frame-parameters", ec.modifyFrameParameters_autogen, 2)
	ec.defSubr1("frame-char-height", ec.frameCharHeight_autogen, 0)
	ec.defSubr1("frame-char-width", ec.frameCharWidth_autogen, 0)
	ec.defSubr1("frame-native-width", ec.frameNativeWidth_autogen, 0)
	ec.defSubr1("frame-native-height", ec.frameNativeHeight_autogen, 0)
	ec.defSubr1("tool-bar-pixel-width", ec.toolBarPixelWidth_autogen, 0)
	ec.defSubr1("frame-text-cols", ec.frameTextCols_autogen, 0)
	ec.defSubr1("frame-text-lines", ec.frameTextLines_autogen, 0)
	ec.defSubr1("frame-total-cols", ec.frameTotalCols_autogen, 0)
	ec.defSubr1("frame-total-lines", ec.frameTotalLines_autogen, 0)
	ec.defSubr1("frame-text-width", ec.frameTextWidth_autogen, 0)
	ec.defSubr1("frame-text-height", ec.frameTextHeight_autogen, 0)
	ec.defSubr1("frame-scroll-bar-width", ec.scrollBarWidth_autogen, 0)
	ec.defSubr1("frame-scroll-bar-height", ec.scrollBarHeight_autogen, 0)
	ec.defSubr1("frame-fringe-width", ec.fringeWidth_autogen, 0)
	ec.defSubr1("frame-child-frame-border-width", ec.frameChildFrameBorderWidth_autogen, 0)
	ec.defSubr1("frame-internal-border-width", ec.frameInternalBorderWidth_autogen, 0)
	ec.defSubr1("frame-right-divider-width", ec.rightDividerWidth_autogen, 0)
	ec.defSubr1("frame-bottom-divider-width", ec.bottomDividerWidth_autogen, 0)
	ec.defSubr4("set-frame-height", ec.setFrameHeight_autogen, 2)
	ec.defSubr4("set-frame-width", ec.setFrameWidth_autogen, 2)
	ec.defSubr4("set-frame-size", ec.setFrameSize_autogen, 3)
	ec.defSubr1("frame-position", ec.framePosition_autogen, 0)
	ec.defSubr3("set-frame-position", ec.setFramePosition_autogen, 3)
	ec.defSubr1("frame-window-state-change", ec.frameWindowStateChange_autogen, 0)
	ec.defSubr2("set-frame-window-state-change", ec.setFrameWindowStateChange_autogen, 0)
	ec.defSubr1("frame-scale-factor", ec.frameScaleFactor_autogen, 0)
	ec.defSubr4("x-get-resource", ec.xGetResource_autogen, 2)
	ec.defSubr1("x-parse-geometry", ec.xParseGeometry_autogen, 1)
	ec.defSubr1("frame-pointer-visible-p", ec.framePointerVisibleP_autogen, 0)
	ec.defSubr2("frame--set-was-invisible", ec.frameSetWasInvisible_autogen, 2)
	ec.defSubr1("reconsider-frame-fonts", ec.reconsiderFrameFonts_autogen, 1)
	ec.defSubr2("eq", ec.eq_autogen, 2)
	ec.defSubr1("null", ec.null_autogen, 1)
	ec.defSubr1("type-of", ec.typeOf_autogen, 1)
	ec.defSubr1("consp", ec.consp_autogen, 1)
	ec.defSubr1("atom", ec.atom_autogen, 1)
	ec.defSubr1("listp", ec.listp_autogen, 1)
	ec.defSubr1("nlistp", ec.nlistp_autogen, 1)
	ec.defSubr1("bare-symbol-p", ec.bareSymbolP_autogen, 1)
	ec.defSubr1("symbol-with-pos-p", ec.symbolWithPosP_autogen, 1)
	ec.defSubr1("symbolp", ec.symbolp_autogen, 1)
	ec.defSubr1("keywordp", ec.keywordp_autogen, 1)
	ec.defSubr1("vectorp", ec.vectorp_autogen, 1)
	ec.defSubr1("recordp", ec.recordp_autogen, 1)
	ec.defSubr1("stringp", ec.stringp_autogen, 1)
	ec.defSubr1("multibyte-string-p", ec.multibyteStringP_autogen, 1)
	ec.defSubr1("char-table-p", ec.charTableP_autogen, 1)
	ec.defSubr1("vector-or-char-table-p", ec.vectorOrCharTableP_autogen, 1)
	ec.defSubr1("bool-vector-p", ec.boolVectorP_autogen, 1)
	ec.defSubr1("arrayp", ec.arrayp_autogen, 1)
	ec.defSubr1("sequencep", ec.sequencep_autogen, 1)
	ec.defSubr1("bufferp", ec.bufferp_autogen, 1)
	ec.defSubr1("markerp", ec.markerp_autogen, 1)
	ec.defSubr1("user-ptrp", ec.userPtrp_autogen, 1)
	ec.defSubr1("subrp", ec.subrp_autogen, 1)
	ec.defSubr1("byte-code-function-p", ec.byteCodeFunctionP_autogen, 1)
	ec.defSubr1("module-function-p", ec.moduleFunctionP_autogen, 1)
	ec.defSubr1("char-or-string-p", ec.charOrStringP_autogen, 1)
	ec.defSubr1("integerp", ec.integerp_autogen, 1)
	ec.defSubr1("integer-or-marker-p", ec.integerOrMarkerP_autogen, 1)
	ec.defSubr1("natnump", ec.natnump_autogen, 1)
	ec.defSubr1("numberp", ec.numberp_autogen, 1)
	ec.defSubr1("number-or-marker-p", ec.numberOrMarkerP_autogen, 1)
	ec.defSubr1("floatp", ec.floatp_autogen, 1)
	ec.defSubr1("threadp", ec.threadp_autogen, 1)
	ec.defSubr1("mutexp", ec.mutexp_autogen, 1)
	ec.defSubr1("condition-variable-p", ec.conditionVariableP_autogen, 1)
	ec.defSubr1("car", ec.car_autogen, 1)
	ec.defSubr1("car-safe", ec.carSafe_autogen, 1)
	ec.defSubr1("cdr", ec.cdr_autogen, 1)
	ec.defSubr1("cdr-safe", ec.cdrSafe_autogen, 1)
	ec.defSubr2("setcar", ec.setcar_autogen, 2)
	ec.defSubr2("setcdr", ec.setcdr_autogen, 2)
	ec.defSubr1("boundp", ec.boundp_autogen, 1)
	ec.defSubr1("fboundp", ec.fboundp_autogen, 1)
	ec.defSubr1("makunbound", ec.makunbound_autogen, 1)
	ec.defSubr1("fmakunbound", ec.fmakunbound_autogen, 1)
	ec.defSubr1("symbol-function", ec.symbolFunction_autogen, 1)
	ec.defSubr1("symbol-plist", ec.symbolPlist_autogen, 1)
	ec.defSubr1("symbol-name", ec.symbolName_autogen, 1)
	ec.defSubr1("bare-symbol", ec.bareSymbol_autogen, 1)
	ec.defSubr1("symbol-with-pos-pos", ec.symbolWithPosPos_autogen, 1)
	ec.defSubr1("remove-pos-from-symbol", ec.removePosFromSymbol_autogen, 1)
	ec.defSubr2("position-symbol", ec.positionSymbol_autogen, 2)
	ec.defSubr2("fset", ec.fset_autogen, 2)
	ec.defSubr3("defalias", ec.defalias_autogen, 2)
	ec.defSubr2("setplist", ec.setplist_autogen, 2)
	ec.defSubr1("subr-arity", ec.subrArity_autogen, 1)
	ec.defSubr1("subr-name", ec.subrName_autogen, 1)
	ec.defSubr1("subr-native-elisp-p", ec.subrNativeElispP_autogen, 1)
	ec.defSubr1("subr-native-lambda-list", ec.subrNativeLambdaList_autogen, 1)
	ec.defSubr1("subr-type", ec.subrType_autogen, 1)
	ec.defSubr1("subr-native-comp-unit", ec.subrNativeCompUnit_autogen, 1)
	ec.defSubr1("native-comp-unit-file", ec.nativeCompUnitFile_autogen, 1)
	ec.defSubr2("native-comp-unit-set-file", ec.nativeCompUnitSetFile_autogen, 2)
	ec.defSubr1("interactive-form", ec.interactiveForm_autogen, 1)
	ec.defSubr1("command-modes", ec.commandModes_autogen, 1)
	ec.defSubr1("indirect-variable", ec.indirectVariable_autogen, 1)
	ec.defSubr1("symbol-value", ec.symbolValue_autogen, 1)
	ec.defSubr2("set", ec.set_autogen, 2)
	ec.defSubr2("add-variable-watcher", ec.addVariableWatcher_autogen, 2)
	ec.defSubr2("remove-variable-watcher", ec.removeVariableWatcher_autogen, 2)
	ec.defSubr1("get-variable-watchers", ec.getVariableWatchers_autogen, 1)
	ec.defSubr1("default-boundp", ec.defaultBoundp_autogen, 1)
	ec.defSubr1("default-value", ec.defaultValue_autogen, 1)
	ec.defSubr2("set-default", ec.setDefault_autogen, 2)
	ec.defSubr1("make-variable-buffer-local", ec.makeVariableBufferLocal_autogen, 1)
	ec.defSubr1("make-local-variable", ec.makeLocalVariable_autogen, 1)
	ec.defSubr1("kill-local-variable", ec.killLocalVariable_autogen, 1)
	ec.defSubr2("local-variable-p", ec.localVariableP_autogen, 1)
	ec.defSubr2("local-variable-if-set-p", ec.localVariableIfSetP_autogen, 1)
	ec.defSubr1("variable-binding-locus", ec.variableBindingLocus_autogen, 1)
	ec.defSubr2("indirect-function", ec.indirectFunction_autogen, 1)
	ec.defSubr2("aref", ec.aref_autogen, 2)
	ec.defSubr3("aset", ec.aset_autogen, 3)
	ec.defSubrM("=", ec.eqlsign_autogen, 1)
	ec.defSubrM("<", ec.lss_autogen, 1)
	ec.defSubrM(">", ec.gtr_autogen, 1)
	ec.defSubrM("<=", ec.leq_autogen, 1)
	ec.defSubrM(">=", ec.geq_autogen, 1)
	ec.defSubr2("/=", ec.neq_autogen, 2)
	ec.defSubr1("number-to-string", ec.numberToString_autogen, 1)
	ec.defSubr2("string-to-number", ec.stringToNumber_autogen, 1)
	ec.defSubrM("+", ec.plus_autogen, 0)
	ec.defSubrM("-", ec.minus_autogen, 0)
	ec.defSubrM("*", ec.times_autogen, 0)
	ec.defSubrM("/", ec.quo_autogen, 1)
	ec.defSubr2("%", ec.rem_autogen, 2)
	ec.defSubr2("mod", ec.mod_autogen, 2)
	ec.defSubrM("max", ec.max_autogen, 1)
	ec.defSubrM("min", ec.min_autogen, 1)
	ec.defSubrM("logand", ec.logand_autogen, 0)
	ec.defSubrM("logior", ec.logior_autogen, 0)
	ec.defSubrM("logxor", ec.logxor_autogen, 0)
	ec.defSubr1("logcount", ec.logcount_autogen, 1)
	ec.defSubr2("ash", ec.ash_autogen, 2)
	ec.defSubr1("1+", ec.add1_autogen, 1)
	ec.defSubr1("1-", ec.sub1_autogen, 1)
	ec.defSubr1("lognot", ec.lognot_autogen, 1)
	ec.defSubr0("byteorder", ec.byteorder_autogen)
	ec.defSubr3("bool-vector-exclusive-or", ec.boolVectorExclusiveOr_autogen, 2)
	ec.defSubr3("bool-vector-union", ec.boolVectorUnion_autogen, 2)
	ec.defSubr3("bool-vector-intersection", ec.boolVectorIntersection_autogen, 2)
	ec.defSubr3("bool-vector-set-difference", ec.boolVectorSetDifference_autogen, 2)
	ec.defSubr2("bool-vector-subsetp", ec.boolVectorSubsetp_autogen, 2)
	ec.defSubr2("bool-vector-not", ec.boolVectorNot_autogen, 1)
	ec.defSubr1("bool-vector-count-population", ec.boolVectorCountPopulation_autogen, 1)
	ec.defSubr3("bool-vector-count-consecutive", ec.boolVectorCountConsecutive_autogen, 3)
	ec.defSubr1("profiler-cpu-start", ec.profilerCpuStart_autogen, 1)
	ec.defSubr0("profiler-cpu-stop", ec.profilerCpuStop_autogen)
	ec.defSubr0("profiler-cpu-running-p", ec.profilerCpuRunningP_autogen)
	ec.defSubr0("profiler-cpu-log", ec.profilerCpuLog_autogen)
	ec.defSubr0("profiler-memory-start", ec.profilerMemoryStart_autogen)
	ec.defSubr0("profiler-memory-stop", ec.profilerMemoryStop_autogen)
	ec.defSubr0("profiler-memory-running-p", ec.profilerMemoryRunningP_autogen)
	ec.defSubr0("profiler-memory-log", ec.profilerMemoryLog_autogen)
	ec.defSubr2("function-equal", ec.functionEqual_autogen, 2)
	ec.defSubr1("lock-file", ec.lockFile_autogen, 1)
	ec.defSubr1("unlock-file", ec.unlockFile_autogen, 1)
	ec.defSubr1("lock-buffer", ec.lockBuffer_autogen, 0)
	ec.defSubr0("unlock-buffer", ec.unlockBuffer_autogen)
	ec.defSubr1("file-locked-p", ec.fileLockedP_autogen, 1)
}

// Subroutines count: 1716
// Constants count: 2
