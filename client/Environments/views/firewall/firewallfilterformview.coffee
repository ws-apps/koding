class FirewallFilterFormView extends KDFormViewWithFields

  constructor: (options = {}, data) ->

    options.fields     =
      type             :
        label          : "Type"
        name           : "type"
        cssClass       : "half"
        itemClass      : KDSelectBox
        selectOptions  : [
          { title      : "IP"       ,  value : "ip"             }
          { title      : "Country"  ,  value : "country"        }
          { title      : "Req./sec" ,  value : "request.second" }
          { title      : "Req./min" ,  value : "request.minute" }
        ]
        validate       :
          rules        :
            required   : yes
          messages     :
            required   : "Please select a filter type"
      value            :
        label          : "Value"
        name           : "value"
        cssClass       : "half"
        placeholder    : "Type a value for your filter..."
        validate       :
          rules        :
            required   : yes
          messages     :
            required   : "Please select a filter type"
      action           :
        label          : "Action"
        name           : "action"
        cssClass       : "half action"
        itemClass      : KDSelectBox
        selectOptions  : [
          { title      : "Allow"            , value : "allow"      }
          { title      : "Block"            , value : "block"      }
          { title      : "Show secure page" , value : "securepage" }
        ]
        validate       :
          rules        :
            required   : yes
          messages     :
            required   : "Please select a action type"
      remove           :
        itemClass      : KDCustomHTMLView
        cssClass       : "delete-button half"
        click          : @bound "destroy"
      enabled          :
        label          : "Enabled"
        cssClass       : "half"
        itemClass      : KodingSwitch
        defaultValue   : yes

    unless options.removable
      options.cssClass = KD.utils.curry "undeletable", options.cssClass

    super options, data

  destroy: ->
    return no  unless @getOptions().removable
    @emit "FirewallFilterRemoved"
    super
